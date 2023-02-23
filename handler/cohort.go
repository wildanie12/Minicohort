package handler

import (
	"log"
	"net/http"

	"github.com/Maintown/Minicohort/database"
	"github.com/Maintown/Minicohort/entity"
	"github.com/Maintown/Minicohort/entity/web"
	"github.com/gin-gonic/gin"
)

// Cohorts Handler struct.
type CohortsHandler struct {
	di *database.DatabaseInstance
}

// NewCohortsHandler creates instance of cohorts handler.
func NewCohortsHandler(di *database.DatabaseInstance) *CohortsHandler {
	return &CohortsHandler{
		di: di,
	}
}

// Index list of cohorts.
func (ch *CohortsHandler) Index(c *gin.Context) {
	data := []entity.Cohort{}

	// Do your gorm database action here kelv...
	_ = ch.di.Mysql.Model(entity.Cohort{}).Find(&data)

	c.JSON(http.StatusOK, gin.H{
		"message": "Get All Cohort",
	})
}

// Show single cohort.
func (ch *CohortsHandler) Show(c *gin.Context) {
	data := entity.Cohort{}
	CohortId := c.Param("CohortId")

	// Do your gorm database action here kelv...
	_ = ch.di.Mysql.Model(entity.Cohort{}).First(&data)

	c.JSON(http.StatusOK, gin.H{
		"ID Cohort": CohortId,
	})
}

// Insert single cohort.
func (ch *CohortsHandler) Insert(c *gin.Context) {
	var cohortInsert web.CohortInsert

	err := c.ShouldBindJSON(&cohortInsert)

	// Do your gorm database action here kelv...

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"Cohort Id":   cohortInsert.CohortId,
		"Cohort Name": cohortInsert.CohortName,
	})
}

// Delete single cohort.
func (ch *CohortsHandler) Delete(c *gin.Context) {

	// Do your gorm database action here kelv...

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Cohort Data From Database",
	})
}
