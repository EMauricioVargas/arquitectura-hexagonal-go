package entities

type Bill struct {
	ID       uint         `json:"id" gorm:"primaryKey"`
	Client   string       `json:"client"`
	Products []BillDetail `gorm:"foreignKey:BillID" json:"products"`
}

type BillDetail struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	BillID    uint    `json:"bill_id"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
	ProductID uint    `json:"product_id"`
	Quantity  uint    `json:"quantity"`
}
