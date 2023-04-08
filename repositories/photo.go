package repositories

import (
	"my-gram/models"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(photo models.Photo) (models.Photo, error)
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{db}
}

func (r *photoRepository) Create(photo models.Photo) (models.Photo, error) {
	err := r.db.Create(&photo).Error

	if err != nil {
		return photo, err
	}

	return photo, nil
}
