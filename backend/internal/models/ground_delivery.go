package models

import (
	"time"

	"gorm.io/gorm"
)

// Struct for maritime deliveries
type GroundDelivery struct {
	gorm.Model

	CustomerID      int `json:"customer_id" validate:"required"`
	ProductQuantity int `json:"product_quantity" validate:"required,gte=1"`

	ProductType  string `json:"product_type" validate:"required" gorm:"type:varchar(100);not null;"`
	WareHouse    string `json:"ware_house" validate:"required" gorm:"type:varchar(100);not null;"`
	VehiclePlate string `json:"vehicle_plate" validate:"required,len=10" gorm:"type:varchar(100);not null;"`
	GuideNumber  string `json:"guide_number" validate:"required,len=10" gorm:"type:varchar(100);not null;"`

	ShippingPrice   float64 `json:"shipping_price" validate:"required,gte=0"`
	DiscountedPrice float64 `json:"discounted_price" validate:"required,gte=0"`

	RegistrationDate time.Time `json:"registration_date" validate:"required,datetime"`
	DeliveryDate     time.Time `json:"delivery_date" validate:"required,datetime"`
}

func (g GroundDelivery) TableName() string {
	return "ground_deliveries"
}
