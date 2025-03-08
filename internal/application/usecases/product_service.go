package usecases

import (
	"errors"
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/domain/entities"
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/domain/repositories"
)

type ProductService interface {
	GetProductByID(id string) (*entities.Product, error)
	GetProducts() ([]entities.Product, error)
}

type productService struct {
	repository repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repository: repo}
}

func (s *productService) GetProductByID(id string) (*entities.Product, error) {
	product, err := s.repository.FindByID(id)
	if err != nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}
func (s *productService) GetProducts() ([]entities.Product, error) {
	products, err := s.repository.FindAll()
	if err != nil {
		return nil, errors.New("products not found")
	}
	return products, nil
}
