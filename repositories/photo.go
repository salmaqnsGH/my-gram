package repositories

import (
	"my-gram/models"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(photo models.Photo) (models.Photo, error)
	FindAll() ([]models.Photo, error)
	FindByUserID(userID uint) ([]models.Photo, error)
	Update(photo models.Photo) (models.Photo, error)
	FindByID(ID uint) (models.Photo, error)
	Delete(ID uint) error
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

func (r *photoRepository) FindAll() ([]models.Photo, error) {
	var photos []models.Photo
	err := r.db.Find(&photos).Error

	if err != nil {
		return photos, err
	}

	return photos, nil
}

func (r *photoRepository) FindByUserID(userID uint) ([]models.Photo, error) {
	var photos []models.Photo
	err := r.db.Where("user_id = ?", userID).Find(&photos).Error

	if err != nil {
		return photos, err
	}

	return photos, nil
}

func (r *photoRepository) Update(photo models.Photo) (models.Photo, error) {
	err := r.db.Save(&photo).Error

	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *photoRepository) FindByID(ID uint) (models.Photo, error) {
	var photo models.Photo
	err := r.db.Where("id = ?", ID).First(&photo).Error

	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *photoRepository) Delete(ID uint) error {
	var photo models.Photo
	if err := r.db.Where("id = ?", ID).First(&photo).Delete(&photo).Error; err != nil {
		return err
	}

	return nil
}
