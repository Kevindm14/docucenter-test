package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model

	FirstName string `json:"first_name" gorm:"type:varchar(100);not null;"`
	LastName  string `json:"last_name" gorm:"type:varchar(100);not null;"`
	Email     string `json:"email" gorm:"type:varchar(100);not null;"`
	Phone     string `json:"phone" gorm:"type:varchar(11);not null;"`
	Address   string `json:"address" gorm:"type:varchar(100);not null;"`
}

func (c Customer) TableName() string {
	return "customers"
}
