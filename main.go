package main

import (
	"github.com/gin-gonic/gin"

	"todo/server/tasks"

	"todo/server/common/db"
)

func main() {

	router := gin.Default()

	dbUrl := db.BuildConnectionString()
	dbHandler := db.Init(dbUrl)

	tasks.RegisterRoutes(router, dbHandler)

	router.Run()
}
