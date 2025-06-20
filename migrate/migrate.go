package main

import (
	"github/shivam261/ClinicManagement/initializers"
	"github/shivam261/ClinicManagement/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	initializers.DB.AutoMigrate(
		&models.Patient{},
		&models.Employee{},
	)
}
