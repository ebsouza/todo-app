package tasks

import (
	"slices"

	"github.com/ebsouza/todo-app/internal/common/orm"
	"github.com/google/uuid"
)

const statusDefault string = "CREATED"

var allStatus []string = []string{statusDefault, "IN_PROGRESS", "DONE"}

type Task struct {
	orm.Base
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func (t *Task) AddData(data TaskData) {
	t.Title = data.Title
	t.Description = data.Description
}

type TaskData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func NewTask() *Task {
	task := &Task{}
	task.ID = uuid.New()
	task.Status = statusDefault
	return task
}

func IsValidStatus(status string) bool {
	return slices.Contains(allStatus, status)
}
