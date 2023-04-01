package tasks

import "todo/server/common/orm"

type Task struct {
	orm.Base
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
