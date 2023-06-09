package repositories

import (
	"my-gram/models"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	Create(socialMedia models.SocialMedia) (models.SocialMedia, error)
	FindByID(ID uint) (models.SocialMedia, error)
	Update(ID uint, socialMedia models.SocialMedia) (models.SocialMedia, error)
	FindAll() ([]models.SocialMedia, error)
	Delete(ID uint) error
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

func (r *socialMediaRepository) FindByID(ID uint) (models.SocialMedia, error) {
	var socialMedia models.SocialMedia
	if err := r.db.Where("id = ?", ID).First(&socialMedia).Error; err != nil {
		return socialMedia, err
	}

	return socialMedia, nil
}

func (r *socialMediaRepository) Update(ID uint, socialMedia models.SocialMedia) (models.SocialMedia, error) {
	err := r.db.Model(&models.SocialMedia{}).Where("id=?", ID).Updates(socialMedia).Error

	if err != nil {
		return socialMedia, err
	}

	return socialMedia, nil
}

func (r *socialMediaRepository) FindAll() ([]models.SocialMedia, error) {
	var socialMedias []models.SocialMedia

	err := r.db.Find(&socialMedias).Error

	if err != nil {
		return socialMedias, err
	}

	return socialMedias, nil
}

func (r *socialMediaRepository) Delete(ID uint) error {
	var socialMedia models.SocialMedia
	if err := r.db.Where("id = ?", ID).First(&socialMedia).Delete(&socialMedia).Error; err != nil {
		return err
	}

	return nil
}
