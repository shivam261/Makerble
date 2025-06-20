package controllers

import (
	"github/shivam261/ClinicManagement/initializers"
	"github/shivam261/ClinicManagement/models"

	"github.com/gin-gonic/gin"
)

func AddPatient(c *gin.Context) {
	var patient models.Patient
	// if name and disease fields are required, you can add validation here
	// for example, if you want to ensure that the name is not empty:
	if patient.Name == "" || patient.Disease == "" {
		c.JSON(400, gin.H{"error": "Name and Disease fields are required"})
		return
	}
	if patient.Age < 0 {
		c.JSON(400, gin.H{"error": "Age cannot be negative"})
		return
	}
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	if err := initializers.DB.Create(&patient).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to add patient"})
		return
	}
	c.JSON(201, patient)
}
