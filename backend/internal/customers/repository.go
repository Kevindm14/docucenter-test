package customers

import "github.com/Kevindm14/docucenter-test/internal/models"

type Repository interface {
	GetCustomers() ([]models.Customer, error)
	GetCustomerById(id int) (models.Customer, error)
	CreateCustomer(customer models.Customer) (models.Customer, error)
	UpdateCustomer(customer models.Customer, id int) (models.Customer, error)
	DeleteCustomer(id int) error
}
