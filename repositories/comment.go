package repositories

import (
	"my-gram/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment models.Comment) (models.Comment, error)
	Update(commentID uint, comment models.Comment) (models.Comment, error)
	FindByID(commentID uint) (models.Comment, error)
	FindAll() ([]models.Comment, error)
	FindByPhotoID(photoID uint) ([]models.Comment, error)
	Delete(commentID uint) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{db}
}

func (r *commentRepository) Create(comment models.Comment) (models.Comment, error) {
	err := r.db.Create(&comment).Error

	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *commentRepository) Update(commentID uint, comment models.Comment) (models.Comment, error) {
	err := r.db.Model(&models.Comment{}).Where("id=?", commentID).Updates(comment).Error

	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *commentRepository) FindByID(commentID uint) (models.Comment, error) {
	var comment models.Comment
	if err := r.db.Where("id = ?", commentID).First(&comment).Error; err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *commentRepository) FindAll() ([]models.Comment, error) {
	var comments []models.Comment

	err := r.db.Find(&comments).Error

	if err != nil {
		return comments, err
	}

	return comments, nil
}

func (r *commentRepository) FindByPhotoID(photoID uint) ([]models.Comment, error) {
	var comments []models.Comment
	if err := r.db.Where("photo_id = ?", photoID).Find(&comments).Error; err != nil {
		return comments, err
	}

	return comments, nil
}

func (r *commentRepository) Delete(commentID uint) error {
	var comment models.Comment
	if err := r.db.Where("id = ?", commentID).First(&comment).Delete(&comment).Error; err != nil {
		return err
	}

	return nil
}
