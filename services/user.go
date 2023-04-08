package services

import (
	"my-gram/models"
	"my-gram/repositories"
)

type UserService interface {
	CreateUser(input models.User) (models.User, error)
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) CreateUser(input models.User) (models.User, error) {
	user := models.User{}

	user.Username = input.Username
	user.Email = input.Email
	user.Password = input.Password // TODO: hash pwd
	user.Age = input.Age

	newUser, err := s.repository.Create(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
