package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	Environment string
	ServiceAddr string
	PG_DSN      string
)

func SetVar() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	Environment = os.Getenv("ENVIRONMENT")
	ServiceAddr = os.Getenv("SERVICE_ADDR")
	PG_DSN = os.Getenv("PG_DSN")

	return nil
}
