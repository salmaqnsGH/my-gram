package services

import (
	"my-gram/models"
	"my-gram/repositories"
)

type SocialMediaService interface {
	CreateSocialMedia(input models.CreateSocialMediaInput) (models.SocialMedia, error)
	GetSocialMediaByID(ID uint) (models.SocialMedia, error)
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
