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
)

type TaskPostRequest struct {
	title string
	description string
}

func TestCreateTask(t *testing.T) {
	dbHandler, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	repository := NewRepository(dbHandler)
	router := InitializeRouter(repository)

	body, _ := json.Marshal(map[string]string{
        "title": "X",
        "description": "Y",
		"status": "OK",
    })
    payload := bytes.NewBuffer(body)
	req, _ := http.NewRequest("POST", "/tasks/", payload)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var task Task
	json.Unmarshal(w.Body.Bytes(), &task)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "X", task.Title)
	assert.Equal(t, "Y", task.Description)
	assert.Equal(t, "OK", task.Status)
}


//https://github.com/learning-go-book/test_examples/tree/master/solver
//https://gin-gonic.com/docs/testing/

//https://pkg.go.dev/net/http/httptest#ResponseRecorder