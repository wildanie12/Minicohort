package database

import (
	"fmt"
	"log"

	"github.com/Maintown/Minicohort/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DatabaseInstance contains all active instance of database connection.
type DatabaseInstance struct {
	mysql *gorm.DB
}

// New creates instance of `DatabaseInstance`.
func New() *DatabaseInstance {
	return &DatabaseInstance{}
}

// ConnectMySQL establishes connection to mysql driver via gorm.
func (di *DatabaseInstance) ConnectMySQL() {
	dsn := "root:password@tcp(127.0.0.1:3306)/Minicohort?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection Error !")
	}

	di.mysql = db
}

// MigrateMySQL creates database structure based on entities provided.
func (di *DatabaseInstance) MigrateMySQL() {
	// put all of your entity migration here ...
	di.mysql.AutoMigrate(&entity.Cohort{})
}

// Close will ends all active connection object.
func (di *DatabaseInstance) Close() {
	fmt.Println("Graceful shutdown...")
	di.Close()
}
