package repository

import (
	groundDeliveries "github.com/Kevindm14/docucenter-test/internal/ground_deliveries"
	"github.com/Kevindm14/docucenter-test/internal/models"
	"gorm.io/gorm"
)

type pgRepository struct {
	DB *gorm.DB
}

func NewPgRepository(db *gorm.DB) groundDeliveries.Repository {
	return &pgRepository{
		DB: db,
	}
}

// ListGroundDeliveries implements ground_deliveries.Repository.
func (pg *pgRepository) ListGroundDeliveries(filter groundDeliveries.Filter) ([]models.GroundDelivery, error) {
	var groundDeliveries []models.GroundDelivery
	err := pg.DB.Scopes(filter.FilterDeliveries).
		Find(&groundDeliveries).Error

		if err != nil {
		return nil, err
	}

	return groundDeliveries, nil
}

// GetGroundDelivery implements ground_deliveries.Repository.
func (pg *pgRepository) GetGroundDeliveryById(id int) (models.GroundDelivery, error) {
	var groundDelivery models.GroundDelivery
	if err := pg.DB.First(&groundDelivery, id).Error; err != nil {
		return models.GroundDelivery{}, err
	}

	return groundDelivery, nil

}

// CreateGroundDelivery implements ground_deliveries.Repository.
func (pg *pgRepository) CreateGroundDelivery(groundDelivery models.GroundDelivery) (models.GroundDelivery, error) {
	if err := pg.DB.Create(&groundDelivery).Error; err != nil {
		return models.GroundDelivery{}, err
	}

	return groundDelivery, nil
}

// DeleteGroundDelivery implements ground_deliveries.Repository.
func (pg *pgRepository) DeleteGroundDelivery(id int) error {
	groundDelivery := models.GroundDelivery{}
	err := pg.DB.Where("id = ?", id).Delete(&groundDelivery).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateGroundDelivery implements ground_deliveries.Repository.
func (pg *pgRepository) UpdateGroundDelivery(groundDelivery models.GroundDelivery) (models.GroundDelivery, error) {
	if err := pg.DB.Save(&groundDelivery).Error; err != nil {
		return models.GroundDelivery{}, err
	}

	return groundDelivery, nil
}
