package tasks

import (
	"todo/server/common/orm"
	"github.com/google/uuid"
)

const StatusDefault = "STARTED"

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
