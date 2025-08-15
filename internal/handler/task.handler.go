package handler

import (
	"net/http"

	"github.com/fmarsico03/resilient-task-service/internal/dto"
	"github.com/fmarsico03/resilient-task-service/internal/httperror"
	"github.com/fmarsico03/resilient-task-service/internal/service"
	"github.com/fmarsico03/resilient-task-service/internal/validations"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service service.TaskService
}

func NewTaskHandler(service service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var request dto.CreateTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		respondWithError(c, httperror.BadRequest("Invalid request body"))
		return
	}

	if err := validations.RequireField(request.Title, "title"); err != nil {
		respondWithError(c, err)
		return
	}

	if err := validations.RequireField(request.Description, "description"); err != nil {
		respondWithError(c, err)
		return
	}

	if err := validations.RequireField(request.ProjectID, "projectId"); err != nil {
		respondWithError(c, err)
		return
	}

	request.UserID = c.GetString("userId")

	task, err := h.service.CreateTask(c, request)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) GetTasksByProjectID(c *gin.Context) {
	projectID := c.Query("projectId")
	userID := c.GetString("userId")
	role := c.GetString("role")

	input := dto.AccessTaskByProjectIDRequest{
		ProjectID: projectID,
		UserID:    userID,
		UserRole:  role,
	}

	if err := validations.ValidateAccessTaskByProjectIDRequest(input); err != nil {
		respondWithError(c, err)
		return
	}

	tasks, err := h.service.GetTasksByProject(c, input)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetTask(c *gin.Context) {
	taskID := c.Param("id")
	userID := c.GetString("userId")
	role := c.GetString("role")

	input := dto.AccessTaskRequest{
		TaskID:   taskID,
		UserID:   userID,
		UserRole: role,
	}

	task, err := h.service.GetTask(c, input)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	taskID := c.Param("id")
	userID := c.GetString("userId")
	role := c.GetString("role")

	input := dto.AccessTaskRequest{
		TaskID:   taskID,
		UserID:   userID,
		UserRole: role,
	}

	err := h.service.DeleteTask(c, input)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	taskID := c.Param("id")
	userID := c.GetString("userId")
	role := c.GetString("role")

	var request dto.UpdateTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		respondWithError(c, httperror.BadRequest("Invalid request body"))
		return
	}

	request.TaskID = taskID
	request.UserID = userID
	request.UserRole = role

	task, err := h.service.UpdateTask(c, request)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, task)
}
