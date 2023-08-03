package http

import (
	"github.com/Kevindm14/docucenter-test/internal/customers"
	"github.com/Kevindm14/docucenter-test/internal/models"
	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	cr customers.Repository
}

// NewCustomerHandler creates a new customer handler
func NewCustomerHandler(cr customers.Repository) *CustomerHandler {
	return &CustomerHandler{
		cr: cr,
	}
}

// GetCustomers gets all customers
func (ch *CustomerHandler) GetCustomers(c *fiber.Ctx) error {
	customers, err := ch.cr.GetCustomers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(customers)
}

// GetCustomerById gets a customer by id
func (ch *CustomerHandler) GetCustomerById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	customer, err := ch.cr.GetCustomerById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(customer)
}

// CreateCustomer creates a customer
func (ch *CustomerHandler) CreateCustomer(c *fiber.Ctx) error {
	customer := models.Customer{}
	if err := c.BodyParser(&customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	customer, err := ch.cr.CreateCustomer(customer)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(customer)
}

// UpdateCustomer updates a customer
func (ch *CustomerHandler) UpdateCustomer(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	customer, err := ch.cr.GetCustomerById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.BodyParser(&customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if customer.ID != uint(id) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Entity ID mismatch",
		})
	}

	customer, err = ch.cr.UpdateCustomer(customer, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(customer)
}

// DeleteCustomer deletes a customer by id
func (ch *CustomerHandler) DeleteCustomer(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = ch.cr.DeleteCustomer(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
