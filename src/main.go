package main

import (
	"todo-api/src/database"
	"todo-api/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
