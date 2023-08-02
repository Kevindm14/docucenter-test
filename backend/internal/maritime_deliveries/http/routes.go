package http

import "github.com/gofiber/fiber/v2"

func SetMaritimeDeliveries(maritimeDeliveriesGroup fiber.Router, m *MaritimeDeliveryHandler) {
	maritimeDeliveriesGroup.Post("/", nil)
	maritimeDeliveriesGroup.Put("/:id", nil)
}
