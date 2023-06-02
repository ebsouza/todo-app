package main

import (
	"github.com/gin-gonic/gin"

	"github.com/ebsouza/todo-app/tasks"

	"github.com/ebsouza/todo-app/common/db"
)


func main() {

	dbUrl := db.BuildConnectionString()
	dbHandler := db.Init(dbUrl)

	repository := tasks.NewRepository(dbHandler)
	router := gin.Default()
	tasks.AddRouterGroup(router, repository)

	router.Run()
}
