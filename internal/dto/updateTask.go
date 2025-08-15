package dto

type TaskStatus string

const (
	TaskStatusEnd     TaskStatus = "end"
	TaskStatusRestart TaskStatus = "restart"
)

type UpdateTaskRequest struct {
	TaskID   string     `json:"task_id"`
	Status   TaskStatus `json:"status" binding:"required"`
	UserID   string     `json:"user_id"`
	UserRole string     `json:"user_role"`
}
