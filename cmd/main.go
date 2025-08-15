package main

import (
	"context"
	"log"

	"github.com/fmarsico03/resilient-task-service/internal/database"
	"github.com/fmarsico03/resilient-task-service/internal/handler"
	"github.com/fmarsico03/resilient-task-service/internal/repository"
	"github.com/fmarsico03/resilient-task-service/internal/router"
	"github.com/fmarsico03/resilient-task-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	r := gin.Default()
	client, db := database.NewMongoDB()
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Printf("Error disconnecting MongoDB: %v", err)
		}
	}()
	log.Printf("Connected to MongoDB: %s", db.Name())

	taskRepo := repository.NewMongoTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)
	router.RegisterTaskRoutes(r, taskHandler)

	log.Println("Starting task-service on :8082")
	if err := r.Run(":8082"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
