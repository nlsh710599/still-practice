package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	Environment string
	ServiceAddr string
	DSN         string
)

func SetVar() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	Environment = os.Getenv("ENVIRONMENT")
	ServiceAddr = os.Getenv("SERVICE_ADDR")
	DSN = os.Getenv("DSN")

	return nil
}
