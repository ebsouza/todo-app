package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func (h handler) PostTasks(ctx *gin.Context) {
	var schema Task

	if err := ctx.BindJSON(&schema); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var task Task
	task.ID = schema.ID
	task.Title = schema.Title
	task.Description = schema.Description
	task.Status = schema.Status

	if result := h.DB.Create(&task); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.IndentedJSON(http.StatusCreated, task)
}

func (h handler) GetTasks(ctx *gin.Context) {
	var tasks []Task

	if result := h.DB.Find(&tasks); result.Error != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, tasks)
}

func (h handler) GetTaskByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var task Task

	if result := h.DB.First(&task, id); result.Error != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, task)

}

func (h handler) RemoveTaskByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var task Task

	if result := h.DB.First(&task, id); result.Error != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	h.DB.Delete(&task)

	ctx.IndentedJSON(http.StatusOK, task)
}
