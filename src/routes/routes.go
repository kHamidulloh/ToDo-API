package routes

import (
	"todo-api/src/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/login", handlers.Login)
	router.POST("/register", handlers.Register)
	authorized := router.Group("/")
	authorized.Use(handlers.AuthMiddleware())
	{
		authorized.POST("/tasks", handlers.CreateTask)
		authorized.GET("/tasks", handlers.GetTasks)
		authorized.PUT("/tasks/:id", handlers.UpdateTask)
		authorized.DELETE("/tasks/:id", handlers.DeleteTask)
	}
}
