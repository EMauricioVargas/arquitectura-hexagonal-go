package db

import (
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/config"
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/domain/entities"
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/domain/repositories"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository() repositories.ProductRepository {
	return &productRepository{db: config.DB}
}

func (r *productRepository) FindByID(id string) (*entities.Product, error) {
	var product entities.Product
	if err := r.db.First(&product, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Create(product *entities.Product) error {
	return r.db.Create(product).Error
}
func (r *productRepository) FindAll() ([]entities.Product, error) {
	var products []entities.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
