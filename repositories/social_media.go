package repositories

import (
	"my-gram/models"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	Create(socialMedia models.SocialMedia) (models.SocialMedia, error)
}

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *socialMediaRepository {
	return &socialMediaRepository{db}
}

func (r *socialMediaRepository) Create(socialMedia models.SocialMedia) (models.SocialMedia, error) {
	err := r.db.Create(&socialMedia).Error

	if err != nil {
		return socialMedia, err
	}

	return socialMedia, nil
}
