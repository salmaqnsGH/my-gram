package repositories

import (
	"my-gram/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment models.Comment) (models.Comment, error)
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
