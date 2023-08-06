package config

import (
	"github.com/voltgizerz/rest-api-go-firestore/logger"

	"github.com/joho/godotenv"
)

// LoadENV - load env file.
func LoadENV() {
	if err := godotenv.Load(); err != nil {
		logger.Log.Warn("No .env file found")
	}
}
