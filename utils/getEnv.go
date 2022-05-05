package utils

import (
	"fmt"

	"github.com/joho/godotenv"
)

// LoadEnv() loads the .env file and returns an error if it fails
func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("Failed to load .env %v", err)
	}
}
