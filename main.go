package main

import (
	"github.com/gin-gonic/gin"

	handlers "todo/server/router"
)


func main() {
	router := gin.Default()
	router.GET("/tasks", handlers.GetTasks)
	router.GET("/tasks/:id", handlers.GetTaskByID)
	router.POST("/tasks", handlers.PostTasks)
	router.DELETE("/tasks/:id", handlers.RemoveTaskByID)

	router.Run(":8080")
}
