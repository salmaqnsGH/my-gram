package services

import (
	"my-gram/models"
	"my-gram/repositories"
)

type PhotoService interface {
	CreatePhoto(input models.Photo) (models.Photo, error)
	GetPhotos() ([]models.Photo, error)
}

type photoService struct {
	repository repositories.PhotoRepository
}

func NewPhotoService(repository repositories.PhotoRepository) *photoService {
	return &photoService{repository}
}

func (s *photoService) CreatePhoto(input models.Photo) (models.Photo, error) {
	photo := models.Photo{}

	// TODO : userID from auth
	photo.UserID = input.UserID
	photo.Title = input.Title
	photo.PhotoUrl = input.PhotoUrl
	photo.Caption = input.Caption

	newPhoto, err := s.repository.Create(photo)
	if err != nil {
		return newPhoto, err
	}

	return newPhoto, nil
}

func (s *photoService) GetPhotos() ([]models.Photo, error) {
	photos, err := s.repository.FindAll()
	if err != nil {
		return photos, err
	}

	return photos, nil
}
