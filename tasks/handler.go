package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	Repository *repository
}

func (h handler) PostTask(ctx *gin.Context) {
	data := TaskData{}

	if err := ctx.BindJSON(&data); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := NewTask()
	task.AddData(data)

	_, err := h.Repository.AddTask(task)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, task)
}

func (h handler) GetTasks(ctx *gin.Context) {
	limit, offset := 100, 0
	
	tasks, err := h.Repository.GetAllTasks(limit, offset)

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
	data := &TaskData{}

	if err := ctx.BindJSON(data); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.Repository.UpdateTask(id, data)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, task)
}
