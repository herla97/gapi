package db

import (
	"fmt"
	"os"

	"github.com/herla97/gapi/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

// Connect : Database connect
func Connect() *gorm.DB {
	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	db, err := gorm.Open("postgres", dbURL)
	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}
