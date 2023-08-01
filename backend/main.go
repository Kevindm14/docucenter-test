package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	log.Println("Starting server on port 8080...")

	app := fiber.New()
	app.Use(logger.New())


	log.Fatal(app.Listen(":8080"))
}
