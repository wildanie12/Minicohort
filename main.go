package main

import(
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/Maintown/Minicohort/Handler"
	"github.com/Maintown/Minicohort/Minicohort"
	"log"
)

func main(){

	dsn := "root:password@tcp(127.0.0.1:3306)/Minicohort?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("Connection Error !")
	}

	db.AutoMigrate(&Minicohort.Cohort{})

	router := gin.Default()
	router.GET("/", Handler.CohortsHandler)
	router.GET("/GetCohort/:CohortId", Handler.GetCohortHandler)
	router.POST("/PostCohortData", Handler.PostCohort)
	router.Run(":8080")
}
