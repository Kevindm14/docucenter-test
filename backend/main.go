package main

import (
	"log"

	"github.com/Kevindm14/docucenter-test/config"
	"github.com/Kevindm14/docucenter-test/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		return c.Next()
	})

	// Database connection
	db := config.PgDBConnection()

	// Routes
	routes.SetRoutesApiV1(app, db)

	log.Fatal(app.Listen(":8080"))
}
