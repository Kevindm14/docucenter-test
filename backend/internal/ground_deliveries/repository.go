package ground_deliveries

import (
	"github.com/Kevindm14/docucenter-test/internal/models"
)

type Repository interface {
	ListGroundDeliveries(filter Filter) ([]models.GroundDelivery, error)
	GetGroundDeliveryById(id int) (models.GroundDelivery, error)
	CreateGroundDelivery(groundDelivery models.GroundDelivery) (models.GroundDelivery, error)
	UpdateGroundDelivery(groundDelivery models.GroundDelivery) (models.GroundDelivery, error)
	DeleteGroundDelivery(id int) error
}
