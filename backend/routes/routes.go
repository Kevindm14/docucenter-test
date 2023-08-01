package routes

import (
	customersHttp "github.com/Kevindm14/docucenter-test/internal/customers/http"
	groundDeliveriesHttp "github.com/Kevindm14/docucenter-test/internal/ground_deliveries/http"
	maritimeDeliveriesHttp "github.com/Kevindm14/docucenter-test/internal/maritime_deliveries/http"
	"github.com/gofiber/fiber/v2"
)

func SetRoutesApiV1(app *fiber.App) {
	api := app.Group("/api/v1")

	// Customers routes
	customersGroup := api.Group("/customers")
	customersHttp.SetCustomerRoutes(customersGroup)

	// Ground deliveries routes
	groundDeliveriesGroup := api.Group("/ground-deliveries")
	groundDeliveriesHttp.SetGroundDeliveries(groundDeliveriesGroup)

	// Maritime deliveries routes
	maritimeDeliveriesGroup := api.Group("/maritime-deliveries")
	maritimeDeliveriesHttp.SetMaritimeDeliveries(maritimeDeliveriesGroup)
}
