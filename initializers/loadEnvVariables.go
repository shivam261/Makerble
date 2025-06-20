package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	// Load environment variables from .env file
	err := godotenv.Load("development.env")
	if err != nil {
		log.Println("no .env file found, using default environment variables")
	}
}
