package tasks

import "github.com/gin-gonic/gin"


func AddRouterGroup(router *gin.Engine, r *repository) {    
    h := &handler{
		Repository: r,
	}

	routes := router.Group("/tasks")
	routes.POST("", h.PostTask)
	routes.GET("", h.GetTasks)
	routes.GET("/:id", h.GetTaskByID)
    routes.PUT("/:id", h.UpdateTask)
	routes.DELETE("/:id", h.RemoveTaskByID)
}
