package services

import (
	"my-gram/models"
	"my-gram/repositories"
)

type PhotoService interface {
	CreatePhoto(input models.Photo) (models.Photo, error)
}

type service struct {
	repository repositories.PhotoRepository
}

func NewPhotoService(repository repositories.PhotoRepository) *service {
	return &service{repository}
}

func (s *service) CreatePhoto(input models.Photo) (models.Photo, error) {
	photo := models.Photo{}

	photo.Title = input.Title
	photo.PhotoUrl = input.PhotoUrl
	photo.Caption = input.Caption

	newPhoto, err := s.repository.Create(photo)
	if err != nil {
		return newPhoto, err
	}

	return newPhoto, nil
}
