package router

import (
	"github.com/fmarsico03/resilient-task-service/internal/handler"
	"github.com/fmarsico03/resilient-task-service/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(r *gin.Engine, h *handler.TaskHandler) {
	api := r.Group("/api/tasks")
	api.Use(middleware.AuthMiddleware())

	api.POST("/", h.CreateTask)
	api.GET("/", h.GetTasksByProjectID)
	api.GET("/:id", h.GetTask)
	api.DELETE("/:id", h.DeleteTask)
	api.PUT("/:id", h.UpdateTask)
}
