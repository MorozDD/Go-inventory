package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	DB *gorm.DB
)

func InitDB() (*gorm.DB, error) {
	var err error
	err = godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	s := os.Getenv("POSTGRES_URL")

	DB, err = gorm.Open("postgres", s)
	if err != nil {
		return nil, err
	}

	return DB, nil
}
