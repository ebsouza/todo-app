package tasks

import "todo/server/common/db"

type Task struct {
	db.Base
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
