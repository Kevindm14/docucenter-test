package http

import (
	"github.com/Kevindm14/docucenter-test/internal/auth"
	"github.com/Kevindm14/docucenter-test/internal/models"
	"github.com/Kevindm14/docucenter-test/libraries"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	ar auth.Repository
}

// Initialize AuthHandler
func NewAuthHandler(ar auth.Repository) *AuthHandler {
	return &AuthHandler{
		ar: ar,
	}
}

// Login customer
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	payload := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	customer, err := h.ar.GetCustomerByEmail(payload.Email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Customer not found",
		})
	}

	if customer.CheckPassword(payload.Password) != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid password",
		})
	}

	token, err := libraries.GenerateToken(customer.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

// Register new customer
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	customer := models.Customer{}
	if err := c.BodyParser(&customer); err != nil {
		return err
	}

	verrs := customer.ValidateCustomer()
	if verrs.HasAny() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": verrs.Errors,
		})
	}

	if err := customer.EncryptPassword(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	customer, err := h.ar.RegisterCustomer(customer)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(customer)
}
