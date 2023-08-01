package http

import "github.com/gofiber/fiber/v2"

func SetGroundDeliveries(groundDeliveriesGroup fiber.Router) {
	groundDeliveriesGroup.Get("/")
	groundDeliveriesGroup.Get("/:id")
	groundDeliveriesGroup.Post("/")
	groundDeliveriesGroup.Put("/:id")
	groundDeliveriesGroup.Delete("/:id")
}
