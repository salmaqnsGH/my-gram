package models

import (
	"log"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Photo represents the model for an photo
type Photo struct {
	GORMModel
	UserID   uint   `json:"user_id"`
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" validate:"required"`
}

// CreatePhotoInput represents the model for an createPhotoInput
type CreatePhotoInput struct {
	Title    string `form:"title" validate:"required"`
	Caption  string `form:"caption"`
	PhotoUrl string `form:"photo_url" validate:"required"`
}

// UpdatePhotoInput represents the model for an updatePhotoInput
type UpdatePhotoInput struct {
	Title    string `form:"title" validate:"required"`
	Caption  string `form:"caption"`
	PhotoUrl string `form:"photo_url" validate:"required"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	validate := validator.New()

	err = validate.Struct(p)
	if err != nil {
		log.Print(err)
	}

	return
}
