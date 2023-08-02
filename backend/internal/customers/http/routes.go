package http

import "github.com/gofiber/fiber/v2"

func SetCustomerRoutes(customerGroup fiber.Router, c *CustomerHandler) {
	customerGroup.Get("/", c.GetCustomers)
	customerGroup.Get("/:id", c.GetCustomerById)
	customerGroup.Post("/", c.CreateCustomer)
	customerGroup.Put("/:id", c.UpdateCustomer)
	customerGroup.Delete("/:id", c.DeleteCustomer)
}
