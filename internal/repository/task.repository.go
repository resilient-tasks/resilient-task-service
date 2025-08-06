package repository

import (
	"context"

	"github.com/fmarsico03/resilient-task-service/internal/model"
)

type TaskRepository interface {
	Save(ctx context.Context, task *model.Task) error
	FindByID(ctx context.Context, id string) (*model.Task, error)
	FindByProjectID(ctx context.Context, projectID string) ([]model.Task, error)
	Delete(ctx context.Context, id string, userID string) error
}
