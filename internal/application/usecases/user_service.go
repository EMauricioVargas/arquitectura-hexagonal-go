package usecases

import (
	"errors"
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/domain/entities"
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/domain/repositories"
)

type UserService interface {
	GetUserByID(id string) (*entities.User, error)
	GetUsers() ([]*entities.User, error)
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repository: repo}
}

func (s *userService) GetUserByID(id string) (*entities.User, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
func (s *userService) GetUsers() ([]*entities.User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return nil, errors.New("users not found")
	}
	return users, nil
}
