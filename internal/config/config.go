package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	Environment string
	ServiceAddr string
	PG_HOST     string
	PG_USER     string
	PG_PWD      string
	PG_DB       string
	PG_PORT     string
)

func SetVar() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	Environment = os.Getenv("ENVIRONMENT")
	ServiceAddr = os.Getenv("SERVICE_ADDR")

	PG_HOST = os.Getenv("PG_HOST")
	PG_USER = os.Getenv("PG_USER")
	PG_PWD = os.Getenv("PG_PWD")
	PG_DB = os.Getenv("PG_DB")
	PG_PORT = os.Getenv("PG_PORT")

	return nil
}
