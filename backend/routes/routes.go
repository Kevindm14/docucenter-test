package routes

import (
	"github.com/Kevindm14/docucenter-test/config"
	authHttp "github.com/Kevindm14/docucenter-test/internal/auth/http"
	customersHttp "github.com/Kevindm14/docucenter-test/internal/customers/http"
	groundDeliveriesHttp "github.com/Kevindm14/docucenter-test/internal/ground_deliveries/http"
	maritimeDeliveriesHttp "github.com/Kevindm14/docucenter-test/internal/maritime_deliveries/http"
	"github.com/Kevindm14/docucenter-test/middleware"

	authRepository "github.com/Kevindm14/docucenter-test/internal/auth/repository"
	customerRepository "github.com/Kevindm14/docucenter-test/internal/customers/repository"
	groundDeliveriesRepository "github.com/Kevindm14/docucenter-test/internal/ground_deliveries/repository"
	maritimeDeliveriesRepository "github.com/Kevindm14/docucenter-test/internal/maritime_deliveries/repository"

	"github.com/gofiber/fiber/v2"
)

func SetRoutesApiV1(app *fiber.App) {
	// Database connection
	db := config.PgDBConnection()

	// API group
	api := app.Group("/api/v1")

	// Auth routes
	authGroup := api.Group("/auth")

	// Inject dependencies
	authRepository := authRepository.NewAuthRepository(db)
	authHandlers := authHttp.NewAuthHandler(authRepository)
	authHttp.SetAuthRoutes(authGroup, authHandlers)

	// Customers routes
	customersGroup := api.Group("/customers", middleware.JwtMiddleWare)

	// Inject dependencies
	customerRepository := customerRepository.NewPgRepository(db)
	customerHandlers := customersHttp.NewCustomerHandler(customerRepository)
	customersHttp.SetCustomerRoutes(customersGroup, customerHandlers)

	// Ground deliveries routes
	groundDeliveriesGroup := api.Group("/ground-deliveries", middleware.JwtMiddleWare)

	// Inject dependencies
	groundDeliveriesRepository := groundDeliveriesRepository.NewPgRepository(db)
	groundDeliveriesHandlers := groundDeliveriesHttp.NewGroundDeliveryHandler(groundDeliveriesRepository)
	groundDeliveriesHttp.SetGroundDeliveries(groundDeliveriesGroup, groundDeliveriesHandlers)

	// Maritime deliveries routes
	maritimeDeliveriesGroup := api.Group("/maritime-deliveries", middleware.JwtMiddleWare)

	// Inject dependencies
	maritimeDeliveriesRepository := maritimeDeliveriesRepository.NewPgRepository(db)
	maritimeDeliveriesHandlers := maritimeDeliveriesHttp.NewMaritimeDeliveryHandler(maritimeDeliveriesRepository)
	maritimeDeliveriesHttp.SetMaritimeDeliveries(maritimeDeliveriesGroup, maritimeDeliveriesHandlers)
}
