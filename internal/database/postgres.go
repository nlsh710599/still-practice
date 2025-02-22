package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(host, username, password, databaseName, port string) (*gorm.DB, error) {
	instance, err := createClient(host, username, password, databaseName, port)

	if err != nil {
		fmt.Printf("Failed to create client: %v", err)
		return nil, err
	}

	return instance, nil
}

func createClient(host, username, password, databaseName, port string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		host, username, password, databaseName, port,
	)

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
