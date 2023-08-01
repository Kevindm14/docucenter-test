package http

import "github.com/gofiber/fiber/v2"

func SetCustomerRoutes(customerGroup fiber.Router) {
	customerGroup.Get("/")
	customerGroup.Get("/:id")
	customerGroup.Post("/")
	customerGroup.Put("/:id")
	customerGroup.Delete("/:id")
}
