package tasks

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := router.Group("/tasks")
	routes.POST("/", h.PostTasks)
	routes.GET("/", h.GetTasks)
	routes.GET("/:id", h.GetTaskByID)
    routes.PUT("/:id", h.UpdateTask)
	routes.DELETE("/:id", h.RemoveTaskByID)
}
