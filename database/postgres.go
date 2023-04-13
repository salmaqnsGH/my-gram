package db

import (
	"fmt"

	"my-gram/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Host      string
	Username  string
	Password  string
	Port      int
	Name      string
	DebugMode string
}

const (
	DB_HOST     = "localhost"
	DB_USER     = "root"
	DB_PASSWORD = "secret"
	DB_PORT     = 5432
	DB_NAME     = "my_gram"
	DEBUG_MODE  = false // true/false
)

var (
	db  *gorm.DB
	err error
)

func StartDB(conf *Database) {
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=disable", DB_USER, DB_NAME, DB_PASSWORD, DB_HOST, DB_PORT)

	if conf.Host != "" {
		dsn = fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=disable", conf.Username, conf.Name, conf.Password, conf.Host, conf.Port)
	}

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if DEBUG_MODE || conf.DebugMode == "true" {
		db.Debug().AutoMigrate(models.User{}, models.Comment{}, models.Photo{}, models.SocialMedia{})
	}

	db.AutoMigrate(models.User{}, models.Comment{}, models.Photo{}, models.SocialMedia{})

	fmt.Println("Successfully connected to database!")
}

func GetDB() *gorm.DB {

	return db
}
