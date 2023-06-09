package services

import (
	"errors"
	"my-gram/models"
	"my-gram/repositories"
)

type UserService interface {
	CreateUser(input models.User) (models.User, error)
	GetUserByID(userID uint) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) CreateUser(input models.User) (models.User, error) {
	user, err := s.repository.FindByEmail(input.Email)
	if err == nil {
		return user, errors.New("email already exist")
	}

	user, err = s.repository.FindByUsername(input.Username)
	if err == nil {
		return user, errors.New("username already exist")
	}

	user = models.User{}

	user.Username = input.Username
	user.Email = input.Email
	user.Password = input.Password
	user.Age = input.Age

	newUser, err := s.repository.Create(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *userService) GetUserByID(userID uint) (models.User, error) {
	user, err := s.repository.FindByID(userID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) GetUserByEmail(email string) (models.User, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) GetUserByUsername(username string) (models.User, error) {
	user, err := s.repository.FindByUsername(username)
	if err != nil {
		return user, err
	}

	return user, nil
}
