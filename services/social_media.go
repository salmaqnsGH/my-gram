package services

import (
	"my-gram/models"
	"my-gram/repositories"
)

type SocialMediaService interface {
	CreateSocialMedia(input models.CreateSocialMediaInput) (models.SocialMedia, error)
	GetSocialMediaByID(ID uint) (models.SocialMedia, error)
	UpdateSocialMedia(inputID uint, inputData models.UpdateSocialMediaInput) (models.SocialMedia, error)
	GetSocialMedias() ([]models.SocialMedia, error)
}

type socialMediaService struct {
	repository repositories.SocialMediaRepository
}

func NewSocialMediaService(repository repositories.SocialMediaRepository) *socialMediaService {
	return &socialMediaService{repository}
}

func (s *socialMediaService) CreateSocialMedia(input models.CreateSocialMediaInput) (models.SocialMedia, error) {
	socialMedia := models.SocialMedia{}

	socialMedia.UserID = input.UserID
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

	socialMedia.Name = inputData.Name
	socialMedia.SocialMediaUrl = inputData.SocialMediaUrl

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
