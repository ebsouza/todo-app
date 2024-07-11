package main

import (
	"github.com/gin-gonic/gin"

	"github.com/ebsouza/todo-app/internal/tasks"

	"github.com/ebsouza/todo-app/internal/common/db"
)

func main() {

	dbUrl := db.BuildConnectionString()
	dbHandler := db.Init(dbUrl)

	repository := tasks.NewRepository(dbHandler)
	router := gin.Default()
	tasks.AddRouterGroup(router, repository)

	router.Run()
}
