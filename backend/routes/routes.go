package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"backend/controllers"
	"backend/ws"
)

func RegisterRoutes(app *fiber.App) {
	// WebSocket routes
	app.Use("/ws/:roomCode", ws.WebSocketHandler)
	app.Get("/ws/:roomCode", websocket.New(ws.JoinRoomSocket))

	// API routes
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
	player.Get("/in-room/:roomID", controllers.GetPlayersInRoom);

	chat := api.Group("/chat")
	chat.Post("/send", controllers.SendMessage)
	chat.Get("/history/:roomID", controllers.GetChatHistory)

	vote := api.Group("/vote")
	vote.Post("/submit", controllers.SubmitVote)
	vote.Get("/results/:roomID/:roundNumber", controllers.GetVoteResults)
}