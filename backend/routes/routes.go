package routes

import (
	"github.com/gofiber/fiber/v2"
	"backend/controllers"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Health check route
	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong"})
	})

	room := api.Group("/rooms")
	room.Post("/create", controllers.CreateRoom)
	room.Post("/join", controllers.JoinRoom)
	room.Post("/start", controllers.StartGame)

	player := api.Group("/players")
	player.Post("/reconnect", controllers.ReconnectPlayer)

	chat := api.Group("/chat")
	chat.Post("/send", controllers.SendMessage)
	chat.Get("/history/:roomID", controllers.GetChatHistory)

	vote := api.Group("/vote")
	vote.Post("/submit", controllers.SubmitVote)
	vote.Get("/results/:roomID/:roundNumber", controllers.GetVoteResults)
}