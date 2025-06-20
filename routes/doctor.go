package routes

import (
	"github/shivam261/ClinicManagement/controllers"
	"github/shivam261/ClinicManagement/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterDoctorRoutes(router *gin.Engine) {
	doctorGroup := router.Group("/doctor")
	doctorGroup.Use(middlewares.IsDoctor())
	{
		// get all patients get request
		doctorGroup.GET("/patients", controllers.GetAllPatients)
		doctorGroup.PUT("/patients/:id", controllers.UpdatePatientById)
		// get patient by id
		// update patient details
	}

}
