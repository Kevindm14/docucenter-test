package models

import (
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model

	FirstName string `json:"first_name" gorm:"type:varchar(100);not null;"`
	LastName  string `json:"last_name" gorm:"type:varchar(100);not null;"`
	Password  string `json:"password" gorm:"type:varchar(100);not null;"`
	Email     string `json:"email" gorm:"type:varchar(100);not null;"`
	Phone     string `json:"phone" gorm:"type:varchar(11);not null;"`
	Address   string `json:"address" gorm:"type:varchar(100);not null;"`
}

func (c Customer) TableName() string {
	return "customers"
}

// EncryptPassword encrypt password
func (c *Customer) EncryptPassword() error {
	enc, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	c.Password = string(enc)

	return nil
}

// CheckPassword check password
func (c *Customer) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(c.Password), []byte(password))
}

func (c *Customer) ValidateCustomer() *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.FirstName, Name: "FirstName"},
		&validators.StringIsPresent{Field: c.LastName, Name: "LastName"},
		&validators.StringIsPresent{Field: c.Password, Name: "Password"},
		&validators.StringIsPresent{Field: c.Email, Name: "Email"},
		&validators.StringIsPresent{Field: c.Phone, Name: "Phone"},
		&validators.StringIsPresent{Field: c.Address, Name: "Address"},
	)
}
