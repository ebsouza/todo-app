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

func createTaskPayload(title, description, status string) *bytes.Buffer {
	data := TaskData{Title: title, Description: description, Status: status}

	body, _ := json.Marshal(data)
	payload := bytes.NewBuffer(body)

	return payload
}

func TestPostTask(t *testing.T) {
	router, _ := setupTestRouters()

	title, description, status := "title", "description", "not used"
	payload := createTaskPayload(title, description, status)
	req, _ := http.NewRequest("POST", "/tasks/", payload)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var task Task
	json.Unmarshal(w.Body.Bytes(), &task)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, title, task.Title)
	assert.Equal(t, description, task.Description)
	assert.Equal(t, StatusDefault, task.Status)
}

func TestGetTasks(t *testing.T) {
	router, repository := setupTestRouters()
	var numberOfTasks int = 3

	for i := 0; i < numberOfTasks; i++ {
		task := NewTask()
		repository.AddTask(task)
	}

	req, _ := http.NewRequest("GET", "/tasks/", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var tasks []Task
	json.Unmarshal(w.Body.Bytes(), &tasks)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, numberOfTasks, len(tasks))
}

func TestGetTaskByID(t *testing.T) {
	router, repository := setupTestRouters()

	task := NewTask()
	id, _ := repository.AddTask(task)

	req, _ := http.NewRequest("GET", "/tasks/"+id.String(), nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var taskRecovered Task
	json.Unmarshal(w.Body.Bytes(), &taskRecovered)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, task.ID, id)
	assert.Equal(t, task.ID, taskRecovered.ID)
	assert.Equal(t, task.Title, taskRecovered.Title)
	assert.Equal(t, task.Description, taskRecovered.Description)
	assert.Equal(t, task.Status, taskRecovered.Status)
}

func TestGetTaskByIDNotFound(t *testing.T) {
	router, _ := setupTestRouters()

	req, _ := http.NewRequest("GET", "/tasks/"+"unexistent id", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestRemoveTaskByID(t *testing.T) {
	router, repository := setupTestRouters()

	task := NewTask()
	id, _ := repository.AddTask(task)

	req, _ := http.NewRequest("DELETE", "/tasks/"+id.String(), nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var taskRemoved Task
	json.Unmarshal(w.Body.Bytes(), &taskRemoved)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, task.ID, taskRemoved.ID)
	assert.Equal(t, task.Title, taskRemoved.Title)
	assert.Equal(t, task.Description, taskRemoved.Description)
	assert.Equal(t, task.Status, taskRemoved.Status)
}

func TestRemoveTaskByIDNotFound(t *testing.T) {
	router, _ := setupTestRouters()

	req, _ := http.NewRequest("DELETE", "/tasks/"+"unexistent id", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUpdateTask(t *testing.T) {
	router, repository := setupTestRouters()

	task := NewTask()
	id, _ := repository.AddTask(task)

	title, description, status := "title", "description", "status"
	payload := createTaskPayload(title, description, status)

	req, _ := http.NewRequest("PUT", "/tasks/"+id.String(), payload)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var taskUpdated Task
	json.Unmarshal(w.Body.Bytes(), &taskUpdated)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, task.ID, taskUpdated.ID)
	assert.Equal(t, title, taskUpdated.Title)
	assert.Equal(t, description, taskUpdated.Description)
	assert.Equal(t, status, taskUpdated.Status)
	assert.NotEqual(t, task.Title, taskUpdated.Title)
	assert.NotEqual(t, task.Description, taskUpdated.Description)
	assert.NotEqual(t, task.Status, taskUpdated.Status)
}

func TestUpdateTaskNotFound(t *testing.T) {
	router, _ := setupTestRouters()

	title, description, status := "title", "description", "status"
	payload := createTaskPayload(title, description, status)

	req, _ := http.NewRequest("PUT", "/tasks/"+"unexistent id", payload)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
