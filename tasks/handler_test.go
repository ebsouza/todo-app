package tasks

import (
	"net/http"
	"encoding/json"
	"gorm.io/gorm"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"bytes"
	"github.com/gin-gonic/gin"
)

type TaskPostRequest struct {
	title string
	description string
}


func setupTestRouters() (*gin.Engine, *repository) {
	dbHandler, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	repository := NewRepository(dbHandler)
	router := InitializeRouter(repository)

	return router, repository
}

func TestCreateTask(t *testing.T) {
	router, _ := setupTestRouters()

	data := map[string]string{
        "title": "X",
        "description": "Y",
		"status": "OK",
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


func TestGetTask(t *testing.T) {
	router, repository := setupTestRouters()

	task := Task{Title: "X", Description: "Y", Status: "Z"}

	repository.AddTask(&task)

	req, _ := http.NewRequest("GET", "/tasks/", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	//var task Task
	//json.Unmarshal(w.Body.Bytes(), &task)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "200", w.Body.String())
	//assert.Equal(t, data["title"], task.Title)
	//assert.Equal(t, data["description"], task.Description)
	//assert.Equal(t, data["status"], task.Status)
}


//https://github.com/learning-go-book/test_examples/tree/master/solver
//https://gin-gonic.com/docs/testing/

//https://pkg.go.dev/net/http/httptest#ResponseRecorder