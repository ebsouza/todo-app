package tasks

import (
	"github.com/ebsouza/todo-app/internal/common/orm"
	"github.com/google/uuid"
)

const StatusDefault = "CREATED"

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
	task.Status = StatusDefault
	return task
}

func allStatus() []string {
	return []string{StatusDefault, "IN_PROGRESS", "DONE"}
}

func IsValidStatus(status string) bool {
	allStatus := allStatus()
	for _, validStatus := range allStatus {
		if status == validStatus {
			return true
		}
	}
	return false
}
