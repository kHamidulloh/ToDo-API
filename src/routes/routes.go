package routes

import (
	"todo-api/src/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/tasks", handlers.CreateTask)
}
