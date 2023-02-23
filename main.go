package main

import (
	"github.com/Maintown/Minicohort/database"
	"github.com/Maintown/Minicohort/handler"
	"github.com/Maintown/Minicohort/router"
	"github.com/gin-gonic/gin"
)

func main() {

	// init database instance
	di := database.New()
	di.ConnectMySQL()
	di.MigrateMySQL()

	// init http instance
	g := gin.Default()

	// register dependency
	cohortsHandler := handler.NewCohortsHandler(di)

	// start http instance.
	router.RegisterRoutes(g, cohortsHandler)
	g.Run(":8080")
}
