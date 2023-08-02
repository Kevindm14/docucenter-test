package maritime_deliveries

import "github.com/Kevindm14/docucenter-test/internal/models"

type Repository interface {
	ListMaritimeDeliveries(filter Filter) ([]models.MaritimeDelivery, error)
	GetMaritimeDelivery(id int) (models.MaritimeDelivery, error)
	CreateMaritimeDelivery(maritimeDelivery models.MaritimeDelivery) (models.MaritimeDelivery, error)
	UpdateMaritimeDelivery(maritimeDelivery models.MaritimeDelivery) (models.MaritimeDelivery, error)
	DeleteMaritimeDelivery(id int) error
}
