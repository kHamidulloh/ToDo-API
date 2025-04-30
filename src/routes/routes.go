package routes

import (
	"todo-api/src/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/tasks", handlers.CreateTask)
	router.GET("/tasks", handlers.GetTasks)
	router.PUT("/tasks/:id", handlers.UpdateTask)
	router.DELETE("/tasks/:id", handlers.DeleteTask)
}
