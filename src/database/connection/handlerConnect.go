package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	loadenv "view_lab/src/loadenv"
)

func validateAndGetDBURL() (string, error) {
	envValue := os.Getenv("SELECT_START_DB")
	if envValue == "" {
		return "", errors.New("SELECT_START_DB is not set")
	}

	var dbURL string
	switch envValue {
	case "DEV":
		dbURL = os.Getenv("DB_URL_DEV")
	case "PROD":
		dbURL = os.Getenv("DB_URL_PROD")
	default:
		return "", fmt.Errorf("invalid SELECT_START_DB value: %s (must be DEV or PROD)", envValue)
	}

	if dbURL == "" {
		return "", fmt.Errorf("DB_URL_%s is not set", envValue)
	}

	return dbURL, nil
}

func HandleModelCreation() {
	if DB == nil {
		log.Fatal("Database connection is not initialized")
	}

	createModel := loadenv.LoadCreateModel()
	if createModel == "" {
		log.Fatal("CREATE_MODEL is not set (must be 'true' or 'false')")
	}

	if createModel != "true" && createModel != "false" {
		log.Fatalf("Invalid CREATE_MODEL value: '%s' (must be 'true' or 'false')", createModel)
	}

	if createModel == "false" {
		log.Println("Model creation skipped (CREATE_MODEL=false)")
		return
	}

	log.Println("Creating tables...")
	if err := CreateTables(DB); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}
	log.Println("Tables created successfully")
}
