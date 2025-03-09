package repositories

import "github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/domain/entities"

type BillRepository interface {
	Create(bill *entities.Bill) error
	FindAll() ([]entities.Bill, error)
	FindByID(id string) (*entities.Bill, error)
}
