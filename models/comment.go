package models

import (
	"log"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Comment represents the model for an comment
type Comment struct {
	GORMModel
	UserID  uint   `gorm:"not null" json:"user_id"`
	PhotoID uint   `gorm:"not null" json:"photo_id"`
	Message string `gorm:"not null" json:"message" validate:"required"`
}

// CreateCommentInput represents the model for an createCommentInput
type CreateCommentInput struct {
	PhotoID uint   `json:"photo_id"`
	Message string `json:"message"`
}

// UpdateCommentInput represents the model for an updateCommentInput
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
