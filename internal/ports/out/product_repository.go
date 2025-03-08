package out

import "github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/domain/entities"

type ProductRepository interface {
	FindByID(id string) (*entities.Product, error)
}
