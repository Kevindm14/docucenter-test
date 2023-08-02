package repository

import (
	"github.com/Kevindm14/docucenter-test/internal/maritime_deliveries"
	"github.com/Kevindm14/docucenter-test/internal/models"
	"gorm.io/gorm"
)

type pgRepository struct {
	DB *gorm.DB
}

// DeleteMaritimeDelivery implements maritime_deliveries.Repository.
func (pg *pgRepository) DeleteMaritimeDelivery(id int) error {
	marineDelivery := models.MaritimeDelivery{}
	if err := pg.DB.Where("id = ?", id).Delete(&marineDelivery).Error; err != nil {
		return err
	}

	return nil
}

// GetMaritimeDelivery implements maritime_deliveries.Repository.
func (pg *pgRepository) GetMaritimeDelivery(id int) (models.MaritimeDelivery, error) {
	maritimeDelivery := models.MaritimeDelivery{}
	if err := pg.DB.Where("id = ?", id).First(&maritimeDelivery).Error; err != nil {
		return models.MaritimeDelivery{}, err
	}

	return maritimeDelivery, nil
}

// ListMaritimeDeliveries implements maritime_deliveries.Repository.
func (pg *pgRepository) ListMaritimeDeliveries(filter maritime_deliveries.Filter) ([]models.MaritimeDelivery, error) {
	maritimeDeliveries := []models.MaritimeDelivery{}
	err := pg.DB.Scopes(filter.FilterDeliveries).
		Find(&maritimeDeliveries).Error

	if err != nil {
		return []models.MaritimeDelivery{}, err
	}

	return maritimeDeliveries, nil
}

func NewPgRepository(db *gorm.DB) maritime_deliveries.Repository {
	return &pgRepository{
		DB: db,
	}
}

// CreateMaritimeDelivery implements maritime_deliveries.Repository.
func (pg *pgRepository) CreateMaritimeDelivery(maritimeDelivery models.MaritimeDelivery) (models.MaritimeDelivery, error) {
	if err := pg.DB.Create(&maritimeDelivery).Error; err != nil {
		return models.MaritimeDelivery{}, err
	}

	return maritimeDelivery, nil
}

// UpdateMaritimeDelivery implements maritime_deliveries.Repository.
func (pg *pgRepository) UpdateMaritimeDelivery(maritimeDelivery models.MaritimeDelivery) (models.MaritimeDelivery, error) {
	if err := pg.DB.Save(&maritimeDelivery).Error; err != nil {
		return models.MaritimeDelivery{}, err
	}

	return maritimeDelivery, nil
}
