package models

import (
	"fmt"
	"my-gram/utils"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// User represents the model for an user
type User struct {
	GORMModel
	Username string `gorm:"not null;uniqueIndex" json:"username" validate:"required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" validate:"required,email"`
	Password string `gorm:"not null" json:"password" validate:"required,min=6"`
	Age      uint   `gorm:"not null" json:"age" validate:"required,gt=8"`
}

type RegisterUserResponse struct {
	GORMModel
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      uint   `json:"age"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	validate := validator.New()

	err = validate.Struct(u)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validationError := range validationErrors {
			switch validationError.Field() {
			case "Username":
				return fmt.Errorf("Username is required")
			case "Email":
				return fmt.Errorf("Email is required and must be a valid email address")
			case "Password":
				return fmt.Errorf("Password is required and must be at least 6 characters long")
			case "Age":
				return fmt.Errorf("Age is required and must be greater than 8")
			}
		}
	}

	hashedPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return
	}

	u.Password = hashedPass

	return
}
