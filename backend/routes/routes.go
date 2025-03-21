package routes

import (
	"github.com/gofiber/fiber/v2"
	"backend/routes"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Health check route
	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong"})
	})

	routes.RegisterRoomRoutes(api.Group("/rooms"))
	routes.RegisterPlayerRoutes(api.Group("/players"))
	routes.RegisterChatRoutes(api.Group("/chat"))
	routes.RegisterVoteRoutes(api.Group("/vote"))

}