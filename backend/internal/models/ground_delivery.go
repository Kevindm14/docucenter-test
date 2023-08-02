package models

import (
	"time"

	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"gorm.io/gorm"
)

// Struct for maritime deliveries
type GroundDelivery struct {
	gorm.Model

	CustomerID      int `json:"customer_id"`
	ProductQuantity int `json:"product_quantity"`

	ProductType  string `json:"product_type" gorm:"type:varchar(100);not null;"`
	WareHouse    string `json:"ware_house" gorm:"type:varchar(100);not null;"`
	VehiclePlate string `json:"vehicle_plate" gorm:"type:varchar(100);not null;" validate:"required,regexp=^[A-Z]{3}[0-9]{3}$"`
	GuideNumber  string `json:"guide_number" gorm:"type:varchar(100);not null;"`

	ShippingPrice   float64 `json:"shipping_price"`
	DiscountedPrice float64 `json:"discounted_price"`

	RegistrationDate time.Time `json:"registration_date"`
	DeliveryDate     time.Time `json:"delivery_date"`
}

func (g GroundDelivery) TableName() string {
	return "ground_deliveries"
}

func (g *GroundDelivery) ValidateGroundDelivery() *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: g.ProductType, Name: "ProductType"},
		&validators.StringIsPresent{Field: g.WareHouse, Name: "WareHouse"},
		&validators.StringIsPresent{Field: g.VehiclePlate, Name: "VehiclePlate"},
		&validators.RegexMatch{Field: g.VehiclePlate, Name: "VehiclePlate", Expr: "^[A-Z]{3}[0-9]{3}$"},
		&validators.StringIsPresent{Field: g.GuideNumber, Name: "GuideNumber"},
		&validators.StringIsPresent{Field: g.RegistrationDate.String(), Name: "RegistrationDate"},
		&validators.StringIsPresent{Field: g.DeliveryDate.String(), Name: "DeliveryDate"},
	)
}
