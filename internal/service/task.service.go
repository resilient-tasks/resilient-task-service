package service

import (
	"context"
	"time"

	"github.com/fmarsico03/resilient-task-service/internal/dto"
	"github.com/fmarsico03/resilient-task-service/internal/httperror"
	"github.com/fmarsico03/resilient-task-service/internal/model"
	"github.com/fmarsico03/resilient-task-service/internal/repository"
	"github.com/fmarsico03/resilient-task-service/internal/utils"
	"github.com/fmarsico03/resilient-task-service/internal/validations"
	"go.mongodb.org/mongo-driver/bson"
)

type TaskService interface {
	CreateTask(ctx context.Context, input dto.CreateTaskRequest) (*model.Task, error)
	GetTask(ctx context.Context, input dto.AccessTaskRequest) (*model.Task, error)
	GetTasksByProject(ctx context.Context, input dto.AccessTaskByProjectIDRequest) ([]model.Task, error)
	DeleteTask(ctx context.Context, input dto.AccessTaskRequest) error
	UpdateTask(ctx context.Context, input dto.UpdateTaskRequest) (*model.Task, error)
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(ctx context.Context, input dto.CreateTaskRequest) (*model.Task, error) {
	now := time.Now()
	startDate, err := utils.ParseDate(input.StartDate, "startDate")
	if err != nil {
		return nil, httperror.BadRequest(err.Error())
	}

	estimatedEndDate, err := utils.ParseDate(input.EstimatedEndDate, "estimatedEndDate")
	if err != nil {
		return nil, httperror.BadRequest(err.Error())
	}

	if err := validations.ValidateDateOrder(startDate, estimatedEndDate); err != nil {
		return nil, err
	}

	task := &model.Task{
		Title:            input.Title,
		Description:      input.Description,
		ProjectID:        input.ProjectID,
		UserID:           input.UserID,
		StartDate:        startDate,
		EstimatedEndDate: estimatedEndDate,
		CreatedAt:        now,
	}

	if err := s.repo.Save(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *taskService) GetTask(ctx context.Context, input dto.AccessTaskRequest) (*model.Task, error) {
	task, err := s.repo.FindByID(ctx, input.TaskID)
	if err != nil {
		return nil, err
	}

	if err := validations.ValidateAccessTaskRequest(input.UserID, task.UserID, input.UserRole); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *taskService) GetTasksByProject(ctx context.Context, input dto.AccessTaskByProjectIDRequest) ([]model.Task, error) {
	return s.repo.FindByProjectID(ctx, input.ProjectID)
}

func (s *taskService) DeleteTask(ctx context.Context, input dto.AccessTaskRequest) error {
	task, err := s.repo.FindByID(ctx, input.TaskID)
	if err != nil {
		return err
	}

	if err := validations.ValidateAccessTaskRequest(input.UserID, task.UserID, input.UserRole); err != nil {
		return err
	}

	return s.repo.Delete(ctx, task.ID.Hex(), task.UserID)
}

func (s *taskService) UpdateTask(ctx context.Context, input dto.UpdateTaskRequest) (*model.Task, error) {
	task, err := s.repo.FindByID(ctx, input.TaskID)
	if err != nil {
		return nil, err
	}

	if err := validations.ValidateAccessTaskRequest(input.UserID, task.UserID, input.UserRole); err != nil {
		return nil, err
	}

	if input.Status != dto.TaskStatusEnd && input.Status != dto.TaskStatusRestart {
		return nil, httperror.BadRequest("Invalid status. Must be 'end' or 'restart'")
	}

	if task.RealEndDate != nil && input.Status == dto.TaskStatusEnd {
		return nil, httperror.BadRequest("Task already ended")
	}

	if input.Status == dto.TaskStatusEnd {
		now := time.Now()
		task.RealEndDate = &now
		task.UpdatedAt = now
	}

	if input.Status == dto.TaskStatusRestart {
		task.RealEndDate = nil
		task.UpdatedAt = time.Now()
	}

	update := bson.M{
		"$set": bson.M{
			"realEndDate": task.RealEndDate,
			"updatedAt":   task.UpdatedAt,
		},
	}
	err = s.repo.Update(ctx, task.ID.Hex(), update)
	if err != nil {
		return nil, err
	}

	return task, nil
}
