package db

import (
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/config"
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/domain/entities"
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/domain/repositories"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() repositories.UserRepository {
	return &userRepository{db: config.DB}
}

func (r *userRepository) FindByID(id string) (*entities.User, error) {
	var user entities.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *entities.User) error {
	return r.db.Create(user).Error
}
func (r *userRepository) FindAll() ([]*entities.User, error) {
	var users []*entities.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
