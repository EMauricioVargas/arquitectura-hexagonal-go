package usecases

import (
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/domain/entities"
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/domain/repositories"
	"strconv"
)

type BillService interface {
	CreateBill(bill *entities.Bill) error
	GetBillByID(id string) (*entities.BillResponse, error)
	GetBills() ([]entities.BillResponse, error)
}

type billService struct {
	billRepo    repositories.BillRepository
	productRepo repositories.ProductRepository
}

func NewBillService(repo repositories.BillRepository, productRepository repositories.ProductRepository) BillService {
	return &billService{
		billRepo:    repo,
		productRepo: productRepository,
	}
}

func (s *billService) CreateBill(bill *entities.Bill) error {
	return s.billRepo.Create(bill)
}

func (s *billService) calculateBillResponse(bill *entities.Bill) (*entities.BillResponse, error) {
	var billResponse entities.BillResponse
	billResponse.ID = bill.ID
	billResponse.Client = bill.Client
	totalPrice := 0.0
	for _, detail := range bill.Products {
		product, err := s.productRepo.FindByID(strconv.FormatUint(uint64(detail.ProductID), 10))
		if err != nil {
			return nil, err
		}

		subtotal := float64(detail.Quantity) * product.Price
		totalPrice += subtotal

		billResponse.Products = append(billResponse.Products, entities.BillDetailResponse{
			ProductID: detail.ProductID,
			Quantity:  detail.Quantity,
			Product: entities.ProductDTO{
				ID:    product.ID,
				Name:  product.Name,
				Price: product.Price,
			},
			Subtotal: subtotal,
		})
	}

	billResponse.Total = totalPrice
	return &billResponse, nil
}

func (s *billService) GetBillByID(id string) (*entities.BillResponse, error) {
	bill, err := s.billRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return s.calculateBillResponse(bill)
}

func (s *billService) GetBills() ([]entities.BillResponse, error) {
	bills, err := s.billRepo.FindAll()
	if err != nil {
		return nil, err
	}
	resp := make([]entities.BillResponse, 0)
	for _, bill := range bills {
		billResponse, err := s.calculateBillResponse(&bill)
		if err != nil {
			return nil, err
		}
		resp = append(resp, *billResponse)
	}
	return resp, nil
}
