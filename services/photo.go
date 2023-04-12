package services

import (
	"fmt"
	"my-gram/models"
	"my-gram/repositories"
)

type PhotoService interface {
	CreatePhoto(input models.CreatePhotoInput) (models.Photo, error)
	GetPhotos() ([]models.Photo, error)
	GetPhotosByUserID(userID uint) ([]models.Photo, error)
	UpdatePhoto(ID uint, input models.Photo) (models.Photo, error)
	GetPhotoByID(ID uint) (models.Photo, error)
	DeletePhoto(ID uint) error
}

type photoService struct {
	repository repositories.PhotoRepository
}

func NewPhotoService(repository repositories.PhotoRepository) *photoService {
	return &photoService{repository}
}

func (s *photoService) CreatePhoto(input models.CreatePhotoInput) (models.Photo, error) {
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

func (s *photoService) GetPhotosByUserID(userID uint) ([]models.Photo, error) {
	photos, err := s.repository.FindByUserID(userID)
	if err != nil {
		return photos, err
	}

	return photos, nil
}

func (s *photoService) UpdatePhoto(ID uint, input models.Photo) (models.Photo, error) {
	photo, err := s.repository.FindByID(ID)
	if err != nil {
		return photo, err
	}

	photo.Caption = input.Caption
	photo.Title = input.Title
	photo.UserID = input.UserID
	photo.PhotoUrl = input.PhotoUrl

	updatedPhoto, err := s.repository.Update(photo)
	if err != nil {
		return updatedPhoto, err
	}

	return updatedPhoto, nil
}

func (s *photoService) GetPhotoByID(ID uint) (models.Photo, error) {
	photo, err := s.repository.FindByID(ID)
	if err != nil {
		return photo, err
	}
	fmt.Println("service", photo)
	return photo, nil
}

func (s *photoService) DeletePhoto(ID uint) error {
	err := s.repository.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}
