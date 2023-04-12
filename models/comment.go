package models

import (
	"log"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Comment struct {
	GORMModel
	UserID  uint   `gorm:"not null" json:"user_id"`
	PhotoID uint   `gorm:"not null" json:"photo_id"`
	Message string `gorm:"not null" json:"message" validate:"required"`
}

type CreateCommentInput struct {
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id"`
	Message string `json:"message"`
}

type UpdateCommentInput struct {
	Message string `json:"message"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	validate := validator.New()

	err = validate.Struct(c)
	if err != nil {
		log.Print(err)
	}

	return
}
