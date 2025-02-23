package database

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func IsDuplicatedKeyError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "duplicate key value violates unique")
}

func IsNotFoundError(err error) bool {
	return err != nil && errors.Is(err, gorm.ErrRecordNotFound)
}

func NewPostgres(dsn string) (*gorm.DB, error) {
	instance, err := createClient(dsn)

	if err != nil {
		fmt.Printf("Failed to create client: %v", err)
		return nil, err
	}

	return instance, nil
}

func createClient(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Failed to open connection: %v", err)
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("Failed to get sql.DB: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(80)
	return db, nil
}
