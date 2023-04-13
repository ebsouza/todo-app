package tasks

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type TaskPostRequest struct {
	title       string
	description string
}

func setupTestRouters() (*gin.Engine, *repository) {
	os.Remove("test.db")
	dbHandler, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	repository := NewRepository(dbHandler)
	router := InitializeRouter(repository)

	return router, repository
}

func TestCreateTask(t *testing.T) {
	router, _ := setupTestRouters()

	data := map[string]string{
		"title":       "X",
		"description": "Y",
		"status":      "OK",
	}

	body, _ := json.Marshal(data)
	payload := bytes.NewBuffer(body)
	req, _ := http.NewRequest("POST", "/tasks/", payload)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var task Task
	json.Unmarshal(w.Body.Bytes(), &task)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, data["title"], task.Title)
	assert.Equal(t, data["description"], task.Description)
	assert.Equal(t, data["status"], task.Status)
}

func TestGetTasks(t *testing.T) {
	router, repository := setupTestRouters()
	var numberOfTasks int = 3

	for i := 0; i < numberOfTasks; i++ {
		task := Task{Title: "A", Description: "B", Status: "C"}
		repository.AddTask(&task)
	}

	req, _ := http.NewRequest("GET", "/tasks/", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var tasks []Task
	json.Unmarshal(w.Body.Bytes(), &tasks)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, numberOfTasks, len(tasks))
}

//https://github.com/learning-go-book/test_examples/tree/master/solver
//https://gin-gonic.com/docs/testing/

//https://pkg.go.dev/net/http/httptest#ResponseRecorder
