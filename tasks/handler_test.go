package tasks

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type RepositorySuite struct {
	suite.Suite
	dbHandler  *gorm.DB
	repository *repository
	router     *gin.Engine
}

func (rs *RepositorySuite) SetupSuite() {
	rs.dbHandler, _ = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	rs.repository = NewRepository(rs.dbHandler)
	rs.router = InitializeRouter(rs.repository)
}

func (rs *RepositorySuite) AfterTest(_, _ string) {
	rs.dbHandler.Exec("DELETE FROM tasks")
}

func createTaskPayload(title, description, status string) *bytes.Buffer {
	data := TaskData{Title: title, Description: description, Status: status}

	body, _ := json.Marshal(data)
	payload := bytes.NewBuffer(body)

	return payload
}

func (rs *RepositorySuite) TestPostTask() {
	title, description, status := "title", "description", "not used"
	payload := createTaskPayload(title, description, status)
	req, _ := http.NewRequest("POST", "/tasks", payload)

	w := httptest.NewRecorder()
	rs.router.ServeHTTP(w, req)

	var task Task
	json.Unmarshal(w.Body.Bytes(), &task)

	assert.Equal(rs.T(), http.StatusCreated, w.Code)
	assert.Equal(rs.T(), title, task.Title)
	assert.Equal(rs.T(), description, task.Description)
	assert.Equal(rs.T(), StatusDefault, task.Status)
}

func (rs *RepositorySuite) TestGetTasks() {
	var numberOfTasks int = 3

	for i := 0; i < numberOfTasks; i++ {
		task := NewTask()
		rs.repository.AddTask(task)
	}

	req, _ := http.NewRequest("GET", "/tasks", nil)

	w := httptest.NewRecorder()
	rs.router.ServeHTTP(w, req)

	var tasks []Task
	json.Unmarshal(w.Body.Bytes(), &tasks)

	assert.Equal(rs.T(), http.StatusOK, w.Code)
	assert.Equal(rs.T(), numberOfTasks, len(tasks))
}

func (rs *RepositorySuite) TestGetTasksLimitOffset() {
	var numberOfTasks int = 5

	for i := 0; i < numberOfTasks; i++ {
		task := NewTask()
		rs.repository.AddTask(task)
	}

	limit, offset := "2", "2"
	limit_integer, _ := strconv.Atoi(limit)
	req, _ := http.NewRequest("GET", "/tasks?limit=" + limit + "&offset=" + offset, nil)

	w := httptest.NewRecorder()
	rs.router.ServeHTTP(w, req)

	var tasks []Task
	json.Unmarshal(w.Body.Bytes(), &tasks)

	assert.Equal(rs.T(), http.StatusOK, w.Code)
	assert.Equal(rs.T(), limit_integer, len(tasks))
}

func (rs *RepositorySuite) TestGetTaskByID() {
	task := NewTask()
	id, _ := rs.repository.AddTask(task)

	req, _ := http.NewRequest("GET", "/tasks/"+id.String(), nil)

	w := httptest.NewRecorder()
	rs.router.ServeHTTP(w, req)

	var taskRecovered Task
	json.Unmarshal(w.Body.Bytes(), &taskRecovered)

	assert.Equal(rs.T(), http.StatusOK, w.Code)
	assert.Equal(rs.T(), task.ID, id)
	assert.Equal(rs.T(), task.ID, taskRecovered.ID)
	assert.Equal(rs.T(), task.Title, taskRecovered.Title)
	assert.Equal(rs.T(), task.Description, taskRecovered.Description)
	assert.Equal(rs.T(), task.Status, taskRecovered.Status)
}

func (rs *RepositorySuite) TestGetTaskByIDNotFound() {
	req, _ := http.NewRequest("GET", "/tasks/"+"unexistent id", nil)

	w := httptest.NewRecorder()
	rs.router.ServeHTTP(w, req)

	assert.Equal(rs.T(), http.StatusNotFound, w.Code)
}

func (rs *RepositorySuite) TestRemoveTaskByID() {
	task := NewTask()
	id, _ := rs.repository.AddTask(task)

	req, _ := http.NewRequest("DELETE", "/tasks/"+id.String(), nil)

	w := httptest.NewRecorder()
	rs.router.ServeHTTP(w, req)

	var taskRemoved Task
	json.Unmarshal(w.Body.Bytes(), &taskRemoved)

	assert.Equal(rs.T(), http.StatusOK, w.Code)
	assert.Equal(rs.T(), task.ID, taskRemoved.ID)
	assert.Equal(rs.T(), task.Title, taskRemoved.Title)
	assert.Equal(rs.T(), task.Description, taskRemoved.Description)
	assert.Equal(rs.T(), task.Status, taskRemoved.Status)
}

func (rs *RepositorySuite) TestRemoveTaskByIDNotFound() {
	req, _ := http.NewRequest("DELETE", "/tasks/"+"unexistent id", nil)

	w := httptest.NewRecorder()
	rs.router.ServeHTTP(w, req)

	assert.Equal(rs.T(), http.StatusNotFound, w.Code)
}

func (rs *RepositorySuite) TestUpdateTask() {
	task := NewTask()
	id, _ := rs.repository.AddTask(task)

	title, description, status := "title", "description", allStatus()[1]
	payload := createTaskPayload(title, description, status)

	req, _ := http.NewRequest("PUT", "/tasks/"+id.String(), payload)

	w := httptest.NewRecorder()
	rs.router.ServeHTTP(w, req)

	var taskUpdated Task
	json.Unmarshal(w.Body.Bytes(), &taskUpdated)

	assert.Equal(rs.T(), http.StatusOK, w.Code)
	assert.Equal(rs.T(), task.ID, taskUpdated.ID)
	assert.Equal(rs.T(), title, taskUpdated.Title)
	assert.Equal(rs.T(), description, taskUpdated.Description)
	assert.Equal(rs.T(), status, taskUpdated.Status)
	assert.NotEqual(rs.T(), task.Title, taskUpdated.Title)
	assert.NotEqual(rs.T(), task.Description, taskUpdated.Description)
	assert.NotEqual(rs.T(), task.Status, taskUpdated.Status)
}

func (rs *RepositorySuite) TestUpdateTaskInvalidStatus() {
	task := NewTask()
	id, _ := rs.repository.AddTask(task)

	title, description, status := "title", "description", "status"
	payload := createTaskPayload(title, description, status)

	req, _ := http.NewRequest("PUT", "/tasks/"+id.String(), payload)

	w := httptest.NewRecorder()
	rs.router.ServeHTTP(w, req)

	var taskUpdated Task
	json.Unmarshal(w.Body.Bytes(), &taskUpdated)

	assert.Equal(rs.T(), http.StatusBadRequest, w.Code)
}

func (rs *RepositorySuite) TestUpdateTaskNotFound() {
	title, description, status := "title", "description", StatusDefault
	payload := createTaskPayload(title, description, status)

	req, _ := http.NewRequest("PUT", "/tasks/"+"unexistent id", payload)

	w := httptest.NewRecorder()
	rs.router.ServeHTTP(w, req)

	assert.Equal(rs.T(), http.StatusNotFound, w.Code)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}
