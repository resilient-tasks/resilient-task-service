package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title            string             `bson:"title" json:"title"`
	Description      string             `bson:"description" json:"description"`
	ProjectID        string             `bson:"projectId" json:"projectId"`
	UserID           string             `bson:"userId" json:"userId"`
	StartDate        time.Time          `bson:"startDate" json:"startDate"`
	EstimatedEndDate time.Time          `bson:"estimatedEndDate" json:"estimatedEndDate"`
	RealEndDate      *time.Time         `bson:"realEndDate,omitempty" json:"realEndDate,omitempty"`
	CreatedAt        time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt        time.Time          `bson:"updatedAt" json:"updatedAt"`
}
