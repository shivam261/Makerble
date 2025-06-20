package routes

import (
	"github/shivam261/ClinicManagement/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		// post for doctor or receptionist  login
		authGroup.POST("/register", controllers.Register) // Register endpoint
		authGroup.POST("/login", controllers.Login)
		authGroup.POST("/logout", controllers.Logout) // Logout endpoint

	}
}
