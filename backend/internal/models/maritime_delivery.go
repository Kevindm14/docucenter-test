package models

import (
	"time"

	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"gorm.io/gorm"
)

// Struct for maritime deliveries
type MaritimeDelivery struct {
	gorm.Model

	CustomerID      int `json:"customer_id" validate:"required"`
	ProductQuantity int `json:"product_quantity" validate:"required,gte=1"`

	ProductType  string `json:"product_type" validate:"required"`
	DeliveryPort string `json:"port" validate:"required"`
	VehiclePlate string `json:"vehicle_plate" validate:"required,regexp=^[A-Z]{3}[0-9]{4}[A-Z]{1}$"`
	GuideNumber  string `json:"guide_number" validate:"required,len=10"`

	ShippingPrice   float64 `json:"shipping_price" validate:"required,gte=0"`
	DiscountedPrice float64 `json:"discounted_price" validate:"required,gte=0"`

	RegistrationDate time.Time `json:"registration_date" validate:"required,datetime"`
	DeliveryDate     time.Time `json:"delivery_date" validate:"required,datetime"`
}

func (m MaritimeDelivery) TableName() string {
	return "maritime_deliveries"
}

func (m *MaritimeDelivery) ValidateMaritimeDelivery() *validate.Errors {
	return validate.Validate(
		&validators.IntIsPresent{Field: m.CustomerID, Name: "CustomerID"},
		&validators.IntIsPresent{Field: m.ProductQuantity, Name: "ProductQuantity"},
		&validators.StringLengthInRange{Field: m.ProductType, Name: "ProductType", Min: 1, Max: 100},
		&validators.StringLengthInRange{Field: m.DeliveryPort, Name: "DeliveryPort", Min: 1, Max: 100},
		&validators.StringLengthInRange{Field: m.VehiclePlate, Name: "VehiclePlate", Min: 1, Max: 100},
		&validators.RegexMatch{Field: m.VehiclePlate, Name: "VehiclePlate", Expr: "^[A-Z]{3}[0-9]{4}[A-Z]{1}$"},
		&validators.StringLengthInRange{Field: m.GuideNumber, Name: "GuideNumber", Min: 1, Max: 10},
		&validators.StringIsPresent{Field: m.RegistrationDate.String(), Name: "RegistrationDate"},
		&validators.StringIsPresent{Field: m.DeliveryDate.String(), Name: "DeliveryDate"},
	)
}
