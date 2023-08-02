package http

import "github.com/gofiber/fiber/v2"

func SetGroundDeliveries(groundDeliveriesGroup fiber.Router, g *GroundDeliveryHandler) {
	groundDeliveriesGroup.Post("/", g.CreateGroundDelivery)
	groundDeliveriesGroup.Get("/", g.ListGroundDeliveries)
	groundDeliveriesGroup.Put("/:id", g.UpdateGroundDelivery)
	groundDeliveriesGroup.Delete("/:id", g.DeleteGroundDelivery)
}
