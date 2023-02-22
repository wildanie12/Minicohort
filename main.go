package main

import(
	"github.com/gin-gonic/gin"
	"github.com/Maintown/Minicohort/Handler"
)

func main(){
	router := gin.Default()
	router.GET("/", Handler.CohortsHandler)
	router.GET("/GetCohort/:CohortId", Handler.GetCohortHandler)
	router.POST("/PostCohortData", Handler.PostCohort)
	router.Run(":8080")
}