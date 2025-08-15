package repository

import (
	"context"

	"github.com/fmarsico03/resilient-task-service/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

type TaskRepository interface {
	Save(ctx context.Context, task *model.Task) error
	FindByID(ctx context.Context, id string) (*model.Task, error)
	FindByProjectID(ctx context.Context, projectID string) ([]model.Task, error)
	Delete(ctx context.Context, id string, userID string) error
	Update(ctx context.Context, id string, update bson.M) error
}
