package routes

import (
	"github.com/gofiber/fiber/v2"
	"backend/controllers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// User routes
	api.Get("/users", controllers.GetUsers)
	api.Post("/users", controllers.CreateUser)
}