package http

import "github.com/gofiber/fiber/v2"

func SetMaritimeDeliveries(maritimeDeliveriesGroup fiber.Router) {
	maritimeDeliveriesGroup.Get("/")
	maritimeDeliveriesGroup.Get("/:id")
	maritimeDeliveriesGroup.Post("/")
	maritimeDeliveriesGroup.Put("/:id")
	maritimeDeliveriesGroup.Delete("/:id")
}
