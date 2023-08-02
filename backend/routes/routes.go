package routes

import (
	customersHttp "github.com/Kevindm14/docucenter-test/internal/customers/http"
	groundDeliveriesHttp "github.com/Kevindm14/docucenter-test/internal/ground_deliveries/http"
	maritimeDeliveriesHttp "github.com/Kevindm14/docucenter-test/internal/maritime_deliveries/http"

	customerRepository "github.com/Kevindm14/docucenter-test/internal/customers/repository"
	groundDeliveriesRepository "github.com/Kevindm14/docucenter-test/internal/ground_deliveries/repository"
	maritimeDeliveriesRepository "github.com/Kevindm14/docucenter-test/internal/maritime_deliveries/repository"

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

func SetRoutesApiV1(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api/v1")

	// Customers routes
	customersGroup := api.Group("/customers")

	// Inject dependencies
	customerRepository := customerRepository.NewPgRepository(db)
	customerHandlers := customersHttp.NewCustomerHandler(customerRepository)
	customersHttp.SetCustomerRoutes(customersGroup, customerHandlers)

	// Ground deliveries routes
	groundDeliveriesGroup := api.Group("/ground-deliveries")

	// Inject dependencies
	groundDeliveriesRepository := groundDeliveriesRepository.NewPgRepository(db)
	groundDeliveriesHandlers := groundDeliveriesHttp.NewGroundDeliveryHandler(groundDeliveriesRepository)
	groundDeliveriesHttp.SetGroundDeliveries(groundDeliveriesGroup, groundDeliveriesHandlers)

	// Maritime deliveries routes
	maritimeDeliveriesGroup := api.Group("/maritime-deliveries")

	// Inject dependencies
	maritimeDeliveriesRepository := maritimeDeliveriesRepository.NewPgRepository(db)
	maritimeDeliveriesHandlers := maritimeDeliveriesHttp.NewMaritimeDeliveryHandler(maritimeDeliveriesRepository)
	maritimeDeliveriesHttp.SetMaritimeDeliveries(maritimeDeliveriesGroup, maritimeDeliveriesHandlers)
}
