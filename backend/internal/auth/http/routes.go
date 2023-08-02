package http

import "github.com/gofiber/fiber/v2"

func SetAuthRoutes(authGroup fiber.Router, ah *AuthHandler) {
	authGroup.Post("/login", ah.Login)
	authGroup.Post("/register", ah.Register)
}
