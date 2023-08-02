package http

import (
	groundDeliveries "github.com/Kevindm14/docucenter-test/internal/ground_deliveries"
	"github.com/Kevindm14/docucenter-test/internal/models"
	"github.com/gofiber/fiber/v2"
)

type GroundDeliveryHandler struct {
	cr groundDeliveries.Repository
}

// Initialize GroundDeliveryHandler
func NewGroundDeliveryHandler(cr groundDeliveries.Repository) *GroundDeliveryHandler {
	return &GroundDeliveryHandler{
		cr: cr,
	}
}

// ListGroundDeliveries
func (h *GroundDeliveryHandler) ListGroundDeliveries(c *fiber.Ctx) error {
	groundDeliveries, err := h.cr.ListGroundDeliveries()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(groundDeliveries)
}

// CreateGroundDelivery
func (h *GroundDeliveryHandler) CreateGroundDelivery(c *fiber.Ctx) error {
	groundDelivery := models.GroundDelivery{}
	if err := c.BodyParser(&groundDelivery); err != nil {
		return err
	}

	if groundDelivery.ProductQuantity > 10 {
		shippingPrice := groundDelivery.ShippingPrice
		groundDelivery.DiscountedPrice = shippingPrice - (shippingPrice * 0.05)
	}

	verrs := groundDelivery.ValidateGroundDelivery()
	if verrs.HasAny() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": verrs.Errors,
		})
	}

	groundDelivery, err := h.cr.CreateGroundDelivery(groundDelivery)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(groundDelivery)
}

// UpdateGroundDelivery
func (h *GroundDeliveryHandler) UpdateGroundDelivery(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	groundDelivery, err := h.cr.GetGroundDeliveryById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.BodyParser(&groundDelivery); err != nil {
		return err
	}

	if groundDelivery.ProductQuantity > 10 {
		shippingPrice := groundDelivery.ShippingPrice
		groundDelivery.DiscountedPrice = shippingPrice - (shippingPrice * 0.05)
	}

	verrs := groundDelivery.ValidateGroundDelivery()
	if verrs.HasAny() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": verrs.Errors,
		})
	}

	groundDelivery, err = h.cr.UpdateGroundDelivery(groundDelivery)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(groundDelivery)
}

// DeleteGroundDelivery
func (h *GroundDeliveryHandler) DeleteGroundDelivery(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = h.cr.DeleteGroundDelivery(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
