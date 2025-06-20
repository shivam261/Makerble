package testutils

// tests/testutils/setup.go

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github/shivam261/ClinicManagement/initializers"
	"github/shivam261/ClinicManagement/models"
)

func SetupTestDB() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database")
	}
	initializers.DB = db

	// Auto migrate your models
	db.AutoMigrate(&models.Employee{}, &models.Patient{})
}

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return r
}
