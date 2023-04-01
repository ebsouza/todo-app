package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	Repository *repository
}

func (h handler) PostTasks(ctx *gin.Context) {
	task := &Task{}

	if err := ctx.BindJSON(task); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Repository.AddTask(task)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, *task)
}

func (h handler) GetTasks(ctx *gin.Context) {
	tasks, err := h.Repository.GetAllTasks()

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, tasks)
}

func (h handler) GetTaskByID(ctx *gin.Context) {
	id := ctx.Param("id")

	task, err := h.Repository.GetTask(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, task)
}

func (h handler) RemoveTaskByID(ctx *gin.Context) {
	id := ctx.Param("id")

	task, err := h.Repository.DeleteTask(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, task)
}

func (h handler) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var task_info Task

	if err := ctx.BindJSON(&task_info); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.Repository.UpdateTask(id, task_info)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, task)
}
