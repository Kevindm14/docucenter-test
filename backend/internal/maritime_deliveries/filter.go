package maritime_deliveries

import (
	"gorm.io/gorm"
)

type Filter struct {
	ClientID        string `query:"client_id"`
	ProductType     string `query:"product_type"`
	ProductQuantity int    `query:"quantity"`
}

// Initialize Filter
func NewFilter() *Filter {
	return &Filter{}
}

func (f *Filter) FilterDeliveries(db *gorm.DB) *gorm.DB {
	if f.ClientID == "" {
		return db.Where("client_id = ?", f.ClientID)
	}

	if f.ProductType == "" {
		return db.Where("product_type = ?", f.ProductType)
	}

	if f.ProductQuantity == 0 {
		return db.Where("product_quantity = ?", f.ProductQuantity)
	}

	return nil
}
