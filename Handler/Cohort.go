package Handler

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
)

func CohortsHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message":"Get All Cohort",
	})
}

func GetCohortHandler(c *gin.Context){
	CohortId := c.Param("CohortId")
	c.JSON(http.StatusOK, gin.H{
		"ID Cohort":CohortId,
	})
}

type CohortInsert struct {
	CohortId int
	CohortName string
}

func PostCohort(c *gin.Context){
	var cohortInsert CohortInsert

	err := c.ShouldBindJSON(&cohortInsert)

	if err != nil{
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"Cohort Id":cohortInsert.CohortId,
		"Cohort Name":cohortInsert.CohortName,
	})
}

func DeleteCohort(c *gin.Context){
	CohortId := c.Param("CohortId")
	c.JSON(http.StatusOK, gin.H{
		"message":"Delete Cohort Data From Database",
	})
}