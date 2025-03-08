package repositories

import "github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/domain/entities"

type UserRepository interface {
	FindByID(id string) (*entities.User, error)
	Create(user *entities.User) error
	FindAll() ([]*entities.User, error)
}
