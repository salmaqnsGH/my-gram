package services

import (
	"fmt"
	"my-gram/models"
	"my-gram/repositories"
	"time"
)

type PhotoService interface {
	CreatePhoto(input models.CreatePhotoInput, userID uint) (models.Photo, error)
	GetPhotos() ([]models.Photo, error)
	GetPhotosByUserID(userID uint) ([]models.Photo, error)
	UpdatePhoto(photoID uint, userID uint, input models.UpdatePhotoInput) (models.Photo, error)
	GetPhotoByID(ID uint) (models.Photo, error)
	DeletePhoto(ID uint) error
}

type photoService struct {
	repository repositories.PhotoRepository
}

func NewPhotoService(repository repositories.PhotoRepository) *photoService {
	return &photoService{repository}
}

func (s *photoService) CreatePhoto(input models.CreatePhotoInput, userID uint) (models.Photo, error) {
	photo := models.Photo{}

	photo.UserID = userID
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

func (s *photoService) UpdatePhoto(photoID uint, userID uint, input models.UpdatePhotoInput) (models.Photo, error) {
	photo, err := s.repository.FindByID(photoID)
	if err != nil {
		return photo, err
	}

	inputTime := "2023-04-12T18:30:37.179191+07:00"
	t, err := time.Parse(time.RFC3339Nano, inputTime)
	if err != nil {
		fmt.Println("Failed to parse input time:", err)
		return photo, err
	}

	photo.Caption = input.Caption
	photo.Title = input.Title
	photo.PhotoUrl = input.PhotoUrl
	photo.UserID = userID
	photo.UpdatedAt = t

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
