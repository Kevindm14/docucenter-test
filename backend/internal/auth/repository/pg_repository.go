package repository

import (
	"github.com/Kevindm14/docucenter-test/internal/auth"
	"github.com/Kevindm14/docucenter-test/internal/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

// Initialize AuthRepository
func NewAuthRepository(db *gorm.DB) auth.Repository {
	return &AuthRepository{
		db: db,
	}
}

// GetCustomerByEmail
func (r *AuthRepository) GetCustomerByEmail(email string) (models.Customer, error) {
	customer := models.Customer{}
	if err := r.db.Where("email = ?", email).First(&customer).Error; err != nil {
		return customer, err
	}

	return customer, nil
}

// RegisterCustomer implements auth.Repository.
func (r *AuthRepository) RegisterCustomer(customer models.Customer) (models.Customer, error) {
	if err := r.db.Create(&customer).Error; err != nil {
		return customer, err
	}

	return customer, nil
}
