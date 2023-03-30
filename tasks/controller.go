package tasks

import "github.com/gin-gonic/gin"


func InitializeRouter(r *repository) *gin.Engine{
	router := gin.Default()
    
    h := &handler{
		Repository: r,
	}

	routes := router.Group("/tasks")
	routes.POST("/", h.PostTasks)
	routes.GET("/", h.GetTasks)
	routes.GET("/:id", h.GetTaskByID)
    routes.PUT("/:id", h.UpdateTask)
	routes.DELETE("/:id", h.RemoveTaskByID)

    return router
}
