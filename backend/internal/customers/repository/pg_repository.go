package repository

import (
	"github.com/Kevindm14/docucenter-test/internal/customers"
	"github.com/Kevindm14/docucenter-test/internal/models"
	"gorm.io/gorm"
)

type pgRepository struct {
	DB *gorm.DB
}

func NewPgRepository(db *gorm.DB) customers.Repository {
	return &pgRepository{
		DB: db,
	}
}

// GetCustomerById implements customers.Repository.
func (pg *pgRepository) GetCustomerById(id int) (models.Customer, error) {
	customer := models.Customer{}
	err := pg.DB.First(&customer, id).Error
	if err != nil {
		return models.Customer{}, err
	}

	return customer, nil
}

// GetCustomers implements customers.Repository.
func (pg *pgRepository) GetCustomers() ([]models.Customer, error) {
	customers := []models.Customer{}
	err := pg.DB.Find(&customers).Error
	if err != nil {
		return nil, err
	}

	return customers, nil
}

// CreateCustomer implements customers.Repository.
func (pg *pgRepository) CreateCustomer(customer models.Customer) (models.Customer, error) {
	err := pg.DB.Create(&customer).Error
	if err != nil {
		return models.Customer{}, err
	}

	return customer, nil
}

// UpdateCustomer implements customers.Repository.
func (pg *pgRepository) UpdateCustomer(customer models.Customer, id int) (models.Customer, error) {
	err := pg.DB.Save(&customer).Error
	if err != nil {
		return models.Customer{}, err
	}

	return customer, nil
}

// DeleteCustomer implements customers.Repository.
func (pg *pgRepository) DeleteCustomer(id int) error {
	customer := models.Customer{}
	err := pg.DB.Where("id = ?", id).Delete(&customer).Error
	if err != nil {
		return err
	}

	return nil
}
