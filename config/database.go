package config

import (
	"GoGym/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func SetupDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	connectionStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode= disable", dbHost, dbUser, dbPassword, dbName, dbPort) 

	DB, err = gorm.Open(postgres.Open(connectionStr))
	if err != nil {
		log.Panic("error trying to connect to the database ", err)
	}

	DB.AutoMigrate(&models.User{})
}
