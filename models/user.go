package models

import (
	"log"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	GORMModel
	Username string `gorm:"not null" json:"username" validate:"required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" validate:"required,email"`
	Password string `gorm:"not null" json:"password" validate:"required,min=6"`
	Age      uint   `gorm:"not null" json:"age" validate:"required,gt=8"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	validate := validator.New()

	err = validate.Struct(u)
	if err != nil {
		log.Print(err)
	}

	return
}
