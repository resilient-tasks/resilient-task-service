package repository

import (
	"context"
	"errors"
	"time"

	"github.com/fmarsico03/resilient-task-service/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoTaskRepository struct {
	collection *mongo.Collection
}

func NewMongoTaskRepository(db *mongo.Database) TaskRepository {
	return &mongoTaskRepository{
		collection: db.Collection("tasks"),
	}
}

func (r *mongoTaskRepository) FindByID(ctx context.Context, id string) (*model.Task, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task ID")
	}

	filter := bson.M{"_id": objectID}
	var task model.Task
	err = r.collection.FindOne(ctx, filter).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *mongoTaskRepository) FindByProjectID(ctx context.Context, projectID string) ([]model.Task, error) {
	filter := bson.M{"projectId": projectID}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []model.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *mongoTaskRepository) Save(ctx context.Context, task *model.Task) error {
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	result, err := r.collection.InsertOne(ctx, task)
	if err != nil {
		return err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		task.ID = oid
	}
	return nil
}

func (r *mongoTaskRepository) Delete(ctx context.Context, id string, userID string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task ID")
	}

	filter := bson.M{"_id": objectID, "userId": userID}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("task not found or not owned by user")
	}
	return nil
}

func (r *mongoTaskRepository) Update(ctx context.Context, id string, update bson.M) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task ID")
	}
	filter := bson.M{"_id": objectID}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}
