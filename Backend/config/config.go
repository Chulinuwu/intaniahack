package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, proceeding with default values")
	}
	// Q : is it gonna stop if it fails to load .env file?
}
