package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"strconv"
)

type TaskSchemaIn struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type TaskSchemaOut struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
}

func fromTaskData(task Task) TaskSchemaOut {
	ts := TaskSchemaOut{ID: task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status}
	return ts
}

func fromTaskDataSlice(tasks []Task) []TaskSchemaOut {
	ts := make([]TaskSchemaOut, 0, len(tasks))
	for _, value := range tasks {
		ts = append(ts, fromTaskData(value))
	}
	return ts
}

type handler struct {
	Repository *repository
}

func (h handler) PostTask(ctx *gin.Context) {
	data := TaskSchemaIn{}

	if err := ctx.BindJSON(&data); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := NewTask()
	task.AddData(data.Title, data.Description)

	_, err := h.Repository.AddTask(task)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, fromTaskData(*task))
}

func (h handler) GetTasks(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "100"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	status := ctx.DefaultQuery("status", "")

	tasks, err := h.Repository.GetAllTasks(limit, offset, status)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, fromTaskDataSlice(tasks))
}

func (h handler) GetTaskByID(ctx *gin.Context) {
	id := ctx.Param("id")

	task, err := h.Repository.GetTask(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, fromTaskData(*task))
}

func (h handler) RemoveTaskByID(ctx *gin.Context) {
	id := ctx.Param("id")

	task, err := h.Repository.DeleteTask(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, fromTaskData(*task))
}

func (h handler) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	data := &TaskSchemaIn{}

	if err := ctx.BindJSON(data); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !IsValidStatus(data.Status) {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	task, err := h.Repository.UpdateTask(id, data.Title, data.Description, data.Status)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, fromTaskData(*task))
}
