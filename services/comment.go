package services

import (
	"fmt"
	"my-gram/models"
	"my-gram/repositories"
	"time"
)

type CommentService interface {
	CreateComment(input models.CreateCommentInput) (models.Comment, error)
	UpdateComment(inputID uint, inputData models.UpdateCommentInput) (models.Comment, error)
	GetCommentByID(commentID uint) (models.Comment, error)
	GetComments() ([]models.Comment, error)
	GetCommentsByPhotoID(photoID uint) ([]models.Comment, error)
	DeleteComment(commentID uint) error
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

	inputTime := "2023-04-12T18:30:37.179191+07:00"
	t, err := time.Parse(time.RFC3339Nano, inputTime)
	if err != nil {
		fmt.Println("Failed to parse input time:", err)
		return comment, err
	}

	comment.Message = inputData.Message
	comment.UpdatedAt = t

	updatedComment, err := s.repository.Update(inputID, comment)
	if err != nil {
		return updatedComment, err
	}

	return updatedComment, nil
}

func (s *commentService) GetCommentByID(commentID uint) (models.Comment, error) {
	comment, err := s.repository.FindByID(commentID)
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (s *commentService) GetComments() ([]models.Comment, error) {
	comments, err := s.repository.FindAll()
	if err != nil {
		return comments, err
	}

	return comments, nil
}

func (s *commentService) GetCommentsByPhotoID(photoID uint) ([]models.Comment, error) {
	comments, err := s.repository.FindByPhotoID(photoID)
	if err != nil {
		return comments, err
	}

	return comments, nil
}

func (s *commentService) DeleteComment(commentID uint) error {
	err := s.repository.Delete(commentID)
	if err != nil {
		return err
	}

	return nil
}
