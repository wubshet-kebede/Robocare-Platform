package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB returns a connected GORM instance instead of using a global var
func InitDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Addis_Ababa",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// using config to add sql logger
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), 
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	//uuid extension enabled
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	log.Println(" Successfully connected to PostgreSQL ")
	return db
}
