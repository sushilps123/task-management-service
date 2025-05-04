package routes

import (
	handlers "task-management/internal/app/task/handlers"
	repository "task-management/internal/app/task/repository"
	"task-management/internal/app/task/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes will define and register all the routes for the task management service.
func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	// Initialize the repository, service, and handler
	repo := repository.NewTaskRepository(db)
	service := services.NewTaskService(repo)
	handler := handlers.NewTaskHandler(service)

	// Register the routes
	r.POST("/tasks", handler.CreateTask)
	r.GET("/tasks", handler.ListTasks)
	r.GET("/tasks/:id", handler.GetTask)
	r.PUT("/tasks/:id", handler.UpdateTask)
	r.DELETE("/tasks/:id", handler.DeleteTask)
}
