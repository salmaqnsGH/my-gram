package services

import (
	"my-gram/models"
	"my-gram/repositories"
)

type CommentService interface {
	CreateComment(input models.CreateCommentInput) (models.Comment, error)
}

type commentService struct {
	repository repositories.CommentRepository
}

func NewCommentService(repository repositories.CommentRepository) *commentService {
	return &commentService{repository}
}

func (s *commentService) CreateComment(input models.CreateCommentInput) (models.Comment, error) {
	comment := models.Comment{}

	comment.UserID = input.UserID
	comment.PhotoID = input.PhotoID
	comment.Message = input.Message

	newComment, err := s.repository.Create(comment)
	if err != nil {
		return newComment, err
	}

	return newComment, nil
}
