package routes

import (
	"github/shivam261/ClinicManagement/controllers"
	"github/shivam261/ClinicManagement/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterReceptionistRoutes(router *gin.Engine) {
	receptionistGroup := router.Group("/receptionist")
	receptionistGroup.Use(middlewares.IsReceptionist())
	{
		// patient registration
		receptionistGroup.POST("/registerPatient", controllers.AddPatient) // Add patient endpoint

	}
}
