package routes

import (
	"github.com/gofiber/fiber/v2"
	"backend/controllers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Health check route
	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong"})
	})

	// User routes
	api.Get("/users", controllers.GetUsers)
	api.Post("/users", controllers.CreateUser)
}