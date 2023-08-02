package auth

import "github.com/Kevindm14/docucenter-test/internal/models"

type Repository interface {
	GetCustomerByEmail(email string) (models.Customer, error)
	RegisterCustomer(customer models.Customer) (models.Customer, error)
}
