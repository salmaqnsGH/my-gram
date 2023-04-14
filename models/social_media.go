package models

import (
	"log"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// SocialMedia represents the model for an socialMedia
type SocialMedia struct {
	GORMModel
	UserID         uint   `gorm:"not null" json:"user_id"`
	Name           string `gorm:"not null" json:"name" validate:"required"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" validate:"required"`
}

// CreateSocialMediaInput represents the model for an createSocialMediaInput
type CreateSocialMediaInput struct {
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}

// UpdateSocialMediaInput represents the model for an updateSocialMediaInput
type UpdateSocialMediaInput struct {
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	validate := validator.New()

	err = validate.Struct(s)
	if err != nil {
		log.Print(err)
	}

	return
}
