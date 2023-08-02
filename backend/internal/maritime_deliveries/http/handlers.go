package http

import (
	maritimeDeliveries "github.com/Kevindm14/docucenter-test/internal/maritime_deliveries"
	"github.com/Kevindm14/docucenter-test/internal/models"
	"github.com/gofiber/fiber/v2"
)

type MaritimeDeliveryHandler struct {
	mr maritimeDeliveries.Repository
}

// Initialize MaritimeDeliveryHandler
func NewMaritimeDeliveryHandler(mr maritimeDeliveries.Repository) *MaritimeDeliveryHandler {
	return &MaritimeDeliveryHandler{
		mr: mr,
	}
}

// GetMaritimeDeliveries
func (m *MaritimeDeliveryHandler) GetMaritimeDeliveries(c *fiber.Ctx) error {
	filterDeliveries := maritimeDeliveries.Filter{
		ClientID:        c.Query("client_id"),
		ProductType:     c.Query("type"),
		ProductQuantity: c.QueryInt("quantity"),
	}

	maritimeDeliveries, err := m.mr.ListMaritimeDeliveries(filterDeliveries)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(maritimeDeliveries)
}

// CreateMaritimeDelivery
func (m *MaritimeDeliveryHandler) CreateMaritimeDelivery(c *fiber.Ctx) error {
	maritimeDelivery := models.MaritimeDelivery{}
	if err := c.BodyParser(&maritimeDelivery); err != nil {
		return err
	}

	if maritimeDelivery.ProductQuantity > 10 {
		shippingPrice := maritimeDelivery.ShippingPrice
		maritimeDelivery.DiscountedPrice = shippingPrice - (shippingPrice * 0.03)
	}

	verrs := maritimeDelivery.ValidateMaritimeDelivery()
	if verrs.HasAny() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": verrs.Errors,
		})
	}

	maritimeDelivery, err := m.mr.CreateMaritimeDelivery(maritimeDelivery)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(maritimeDelivery)
}

// UpdateMaritimeDelivery
func (h *MaritimeDeliveryHandler) UpdateMaritimeDelivery(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	maritimeDelivery, err := h.mr.GetMaritimeDelivery(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.BodyParser(&maritimeDelivery); err != nil {
		return err
	}

	if maritimeDelivery.ProductQuantity > 10 {
		shippingPrice := maritimeDelivery.ShippingPrice
		maritimeDelivery.DiscountedPrice = shippingPrice - (shippingPrice * 0.03)
	}

	verrs := maritimeDelivery.ValidateMaritimeDelivery()
	if verrs.HasAny() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": verrs.Errors,
		})
	}

	maritimeDelivery, err = h.mr.UpdateMaritimeDelivery(maritimeDelivery)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(maritimeDelivery)
}

// DeleteGroundDelivery
func (h *MaritimeDeliveryHandler) DeleteMaritimeDelivery(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = h.mr.DeleteMaritimeDelivery(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
