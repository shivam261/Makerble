package controllers

import (
	"github/shivam261/ClinicManagement/initializers"
	"github/shivam261/ClinicManagement/models"

	"github.com/gin-gonic/gin"
)

func GetAllPatients(c *gin.Context) {
	// This function will return all the patients in the database
	var patients []models.Patient
	if err := initializers.DB.Find(&patients).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch patients"})
		return
	}
	c.JSON(200, patients)
}
func UpdatePatientById(c *gin.Context) {
	// This function will update the patient details by id
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	id := c.Param("id")
	if err := initializers.DB.Model(&models.Patient{}).Where("id = ?", id).Updates(patient).Error; err != nil {
		c.JSON(500, gin.H{"error": "Pateint Not Found "})
		return
	}
	c.JSON(200, gin.H{"message": "Patient updated successfully"})
}
