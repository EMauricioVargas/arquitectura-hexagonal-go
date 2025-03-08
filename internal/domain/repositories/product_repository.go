package repositories

import "github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/domain/entities"

type ProductRepository interface {
	FindByID(id string) (*entities.Product, error)
	Create(product *entities.Product) error
	FindAll() ([]entities.Product, error)
}
