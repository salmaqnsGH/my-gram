package db

import (
	"fmt"
	"os"
	"strconv"

	"my-gram/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")
	DEBUG_MODE := os.Getenv("DEBUG_MODE")

	dbPort, err := strconv.Atoi(DB_PORT)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=disable", DB_USER, DB_NAME, DB_PASSWORD, DB_HOST, dbPort)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if DEBUG_MODE == "true" {
		db.Debug().AutoMigrate(models.User{}, models.Comment{}, models.Photo{}, models.SocialMedia{})
	}

	db.AutoMigrate(models.User{}, models.Comment{}, models.Photo{}, models.SocialMedia{})

	fmt.Println("Successfully connected to database!")
}

func GetDB() *gorm.DB {
	return db
}
