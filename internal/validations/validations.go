package validations

import (
	"time"

	"github.com/fmarsico03/resilient-task-service/internal/dto"
	"github.com/fmarsico03/resilient-task-service/internal/httperror"
)

func RequireField(fieldValue string, fieldName string) error {
	if fieldValue == "" {
		return httperror.BadRequest(fieldName + " is required")
	}
	return nil
}

func ValidateOwnership(currentUserId, ownerId string, isAdmin bool) error {
	if currentUserId != ownerId && !isAdmin {
		return httperror.Forbidden("You are not allowed to modify this resource")
	}
	return nil
}

func ValidateDateOrder(startDate, estimatedEndDate time.Time) error {
	if !startDate.Before(estimatedEndDate) {
		return httperror.BadRequest("startDate must be before estimatedEndDate")
	}
	return nil
}

func ValidateAccessTaskByProjectIDRequest(input dto.AccessTaskByProjectIDRequest) error {
	if input.ProjectID == "" {
		return httperror.BadRequest("projectId is required")
	}
	if input.UserID == "" {
		return httperror.BadRequest("userId is required")
	}
	if input.UserRole == "" {
		return httperror.BadRequest("userRole is required")
	}
	return nil
}

func ValidateAccessTaskRequest(currentUserId, targetUserId string, role string) error {
	if currentUserId != targetUserId && role != "admin" {
		return httperror.Forbidden("You are not allowed to access this task")
	}
	return nil
}
