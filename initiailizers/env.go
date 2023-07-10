package initiailizers

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	loadEnvVariables()
}

func loadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
