package tasks

import (
	"github.com/ebsouza/todo-app/internal/common/orm"
	"github.com/google/uuid"
)

type taskStatus int

const (
	Created = iota + 1
	InProgress
	Done
)

var statusName = map[taskStatus]string{
	Created:    "created",
	InProgress: "inProgress",
	Done:       "done",
}

func (ts taskStatus) String() string {
	return statusName[ts]
}

func IsValidStatus(status string) bool {
	for _, value := range statusName {
		if status == value {
			return true
		}
	}
	return false
}

type Task struct {
	orm.Base
	Title       string
	Description string
	Status      string
}

func (t *Task) AddData(title string, description string) {
	t.Title = title
	t.Description = description
}

func NewTask() *Task {
	task := &Task{}
	task.ID = uuid.New()
	task.Status = statusName[Created]
	return task
}
