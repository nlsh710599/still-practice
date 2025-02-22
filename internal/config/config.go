package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	Environment string
	ServiceAddr string
)

func SetVar() error {
	err := godotenv.Load(".env") // Always load .env
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	Environment = os.Getenv("ENVIRONMENT")
	ServiceAddr = os.Getenv("SERVICE_ADDR")

	return nil
}
