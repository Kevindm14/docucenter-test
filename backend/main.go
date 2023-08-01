package main

import (
	"log"

	"github.com/Kevindm14/docucenter-test/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	log.Println("Starting server on port 8080...")

	app := fiber.New()
	app.Use(logger.New())

	config.PgDBConnection()

	log.Fatal(app.Listen(":8080"))
}
