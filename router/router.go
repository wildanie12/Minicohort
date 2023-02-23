package router

import (
	"github.com/Maintown/Minicohort/handler"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes attach endpoints to corresponding handler.
func RegisterRoutes(g *gin.Engine, cohort *handler.CohortsHandler) {
	g.GET("/", cohort.Index)
	g.GET("/GetCohort/:CohortId", cohort.Show)
	g.POST("/PostCohortData", cohort.Insert)
}
