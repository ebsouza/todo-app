package main

import (
	"github.com/ebsouza/todo-app/tasks"

	"github.com/ebsouza/todo-app/common/db"
)

func main() {

	dbUrl := db.BuildConnectionString()
	dbHandler := db.Init(dbUrl)

	repository := tasks.NewRepository(dbHandler)
	router := tasks.InitializeRouter(repository)

	router.Run()
}
