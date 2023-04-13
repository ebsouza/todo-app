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
    })
    payload := bytes.NewBuffer(body)
	req, _ := http.NewRequest("POST", "/tasks/", payload)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	//assert.Equal(t, "pong", w.Body.String())
}


//https://github.com/learning-go-book/test_examples/tree/master/solver
//https://gin-gonic.com/docs/testing/