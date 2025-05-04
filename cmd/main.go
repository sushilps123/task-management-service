package main

import (
	"log"
	models "task-management/internal/app/task/models"
	routes "task-management/internal/app/task/routes"
	"task-management/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database := db.InitDB()
	database.AutoMigrate(&models.Task{})

	routes.RegisterRoutes(r, database)

	log.Println("Server running at http://localhost:8080")
	r.Run(":8080")
}
