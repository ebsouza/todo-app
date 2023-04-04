package main

import (
	"todo/server/tasks"

	"todo/server/common/db"
)

func main() {

	dbUrl := db.BuildConnectionString()
	dbHandler := db.Init(dbUrl)

	repository := tasks.NewRepository(dbHandler)
	router := tasks.InitializeRouter(repository)

	router.Run()
}
