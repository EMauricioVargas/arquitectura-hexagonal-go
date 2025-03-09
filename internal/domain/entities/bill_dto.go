package entities

type BillResponse struct {
	ID       uint                 `json:"id"`
	Client   string               `json:"client"`
	Products []BillDetailResponse `json:"products"`
	Total    float64              `json:"total"`
}

type BillDetailResponse struct {
	ProductID uint       `json:"product_id"`
	Product   ProductDTO `json:"product"`
	Quantity  uint       `json:"quantity"`
	Subtotal  float64    `json:"subtotal"` // Cantidad * Precio
}
type ProductDTO struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
