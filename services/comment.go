package services

import (
	"my-gram/models"
	"my-gram/repositories"
)

type CommentService interface {
	CreateComment(input models.CreateCommentInput) (models.Comment, error)
	UpdateComment(inputID uint, inputData models.UpdateCommentInput) (models.Comment, error)
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

func (s *commentService) UpdateComment(inputID uint, inputData models.UpdateCommentInput) (models.Comment, error) {
	comment, err := s.repository.FindByID(inputID)
	if err != nil {
		return comment, err
	}

	comment.Message = inputData.Message

	updatedComment, err := s.repository.Update(inputID, comment)
	if err != nil {
		return updatedComment, err
	}

	return updatedComment, nil
}