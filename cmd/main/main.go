package main

import (
	"github/shivam261/ClinicManagement/initializers"
	"github/shivam261/ClinicManagement/middlewares"
	"github/shivam261/ClinicManagement/routes"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.ConnectToRedis()
}
func main() {
	// This is the main entry point of the application.
	// You can add your application logic here.
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode) // Set Gin to release mode for production
	} else {
		gin.SetMode(gin.DebugMode) // Set Gin to debug mode for development
	}
	router := gin.Default()
	// Set up middleware
	router.Use(middlewares.RateLimiter(100, 1*time.Minute)) // Limit to 100 requests per minute
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Clinic Management System",
		})
	})
	routes.RegisterAuthRoutes(router)
	routes.RegisterDoctorRoutes(router)
	routes.RegisterReceptionistRoutes(router)
	router.Run(":3000") // Start the server on port 3000

}
