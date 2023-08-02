package http

import "github.com/gofiber/fiber/v2"

func SetMaritimeDeliveries(maritimeDeliveriesGroup fiber.Router, m *MaritimeDeliveryHandler) {
	maritimeDeliveriesGroup.Get("/", m.GetMaritimeDeliveries)
	maritimeDeliveriesGroup.Post("/", m.CreateMaritimeDelivery)
	maritimeDeliveriesGroup.Put("/:id", m.UpdateMaritimeDelivery)
	maritimeDeliveriesGroup.Delete("/:id", m.DeleteMaritimeDelivery)
}
