package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"todo/server/domain"

	"todo/server/repository"
)

func GetTasks(c *gin.Context) {
	all_tasks, _ := repository.GetAllTasks()
	c.IndentedJSON(http.StatusOK, all_tasks)
}

func GetTaskByID(c *gin.Context) {
	task_id := c.Param("id")

	task, err := repository.GetTask(task_id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, task)

}

func PostTasks(c *gin.Context) {
	var newTask domain.Task

	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := repository.AddTask(newTask)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newTask)
}

func RemoveTaskByID(c *gin.Context) {
	task_id := c.Param("id")

	task, err := repository.RemoveTask(task_id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}
