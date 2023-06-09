package services

import (
	"fmt"
	"my-gram/models"
	"my-gram/repositories"
	"time"
)

type SocialMediaService interface {
	CreateSocialMedia(input models.CreateSocialMediaInput, userID uint) (models.SocialMedia, error)
	GetSocialMediaByID(ID uint) (models.SocialMedia, error)
	UpdateSocialMedia(inputID uint, inputData models.UpdateSocialMediaInput) (models.SocialMedia, error)
	GetSocialMedias() ([]models.SocialMedia, error)
	DeleteSocialMedia(ID uint) error
}

type socialMediaService struct {
	repository repositories.SocialMediaRepository
}

func NewSocialMediaService(repository repositories.SocialMediaRepository) *socialMediaService {
	return &socialMediaService{repository}
}

func (s *socialMediaService) CreateSocialMedia(input models.CreateSocialMediaInput, userID uint) (models.SocialMedia, error) {
	socialMedia := models.SocialMedia{}

	socialMedia.UserID = userID
	socialMedia.Name = input.Name
	socialMedia.SocialMediaUrl = input.SocialMediaUrl

	newSocialMedia, err := s.repository.Create(socialMedia)
	if err != nil {
		return newSocialMedia, err
	}

	return newSocialMedia, nil
}

func (s *socialMediaService) GetSocialMediaByID(ID uint) (models.SocialMedia, error) {
	socialMedia, err := s.repository.FindByID(ID)
	if err != nil {
		return socialMedia, err
	}

	return socialMedia, nil
}

func (s *socialMediaService) UpdateSocialMedia(inputID uint, inputData models.UpdateSocialMediaInput) (models.SocialMedia, error) {
	socialMedia, err := s.repository.FindByID(inputID)
	if err != nil {
		return socialMedia, err
	}

	inputTime := "2023-04-12T18:30:37.179191+07:00"
	t, err := time.Parse(time.RFC3339Nano, inputTime)
	if err != nil {
		fmt.Println("Failed to parse input time:", err)
		return socialMedia, err
	}

	socialMedia.Name = inputData.Name
	socialMedia.SocialMediaUrl = inputData.SocialMediaUrl
	socialMedia.UpdatedAt = t

	updatedSocialMedia, err := s.repository.Update(inputID, socialMedia)
	if err != nil {
		return updatedSocialMedia, err
	}

	return updatedSocialMedia, nil
}

func (s *socialMediaService) GetSocialMedias() ([]models.SocialMedia, error) {
	socialMedias, err := s.repository.FindAll()
	if err != nil {
		return socialMedias, err
	}

	return socialMedias, nil
}

func (s *socialMediaService) DeleteSocialMedia(ID uint) error {
	err := s.repository.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}
