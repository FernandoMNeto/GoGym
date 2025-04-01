package config

import (
	"GoGym/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}

	connectionStr := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=" + os.Getenv("DB_SSLMODE")

	DB, err = gorm.Open(postgres.Open(connectionStr))
	if err != nil {
		log.Panic("error trying to connect to the database ", err)
	}

	DB.AutoMigrate(&models.Student{})
}
