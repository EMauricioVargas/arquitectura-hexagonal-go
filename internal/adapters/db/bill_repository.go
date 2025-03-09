package db

import (
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/config"
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/domain/entities"
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/domain/repositories"
	"gorm.io/gorm"
)

type billRepository struct {
	db *gorm.DB
}

func NewBillRepository() repositories.BillRepository {
	return &billRepository{db: config.DB}
}

func (r *billRepository) FindAll() ([]entities.Bill, error) {
	var bills []entities.Bill
	if err := r.db.Preload("Products").Find(&bills).Error; err != nil {
		return nil, err
	}
	return bills, nil
}

func (r *billRepository) FindByID(id string) (*entities.Bill, error) {
	var bill entities.Bill
	if err := r.db.Preload("Products").First(&bill, id).Error; err != nil {
		return nil, err
	}
	return &bill, nil
}

func (r *billRepository) Create(bill *entities.Bill) error {
	return r.db.Create(bill).Error
}
