package models

import (
	"time"

	"gorm.io/gorm"
)

// Struct for maritime deliveries
type MaritimeDelivery struct {
	gorm.Model

	CustomerID      int `json:"customer_id" validate:"required"`
	ProductQuantity int `json:"product_quantity" validate:"required,gte=1"`

	ProductType  string `json:"product_type" validate:"required"`
	DeliveryPort string `json:"port" validate:"required"`
	VehiclePlate string `json:"vehicle_plate" validate:"required,len=10"`
	GuideNumber  string `json:"guide_number" validate:"required,len=10"`

	ShippingPrice   float64 `json:"shipping_price" validate:"required,gte=0"`
	DiscountedPrice float64 `json:"discounted_price" validate:"required,gte=0"`

	RegistrationDate time.Time `json:"registration_date" validate:"required,datetime"`
	DeliveryDate     time.Time `json:"delivery_date" validate:"required,datetime"`
}

func (m MaritimeDelivery) TableName() string {
	return "maritime_deliveries"
}
