package controllers

import (
	"backend/database"
	"backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

// Request payload for creating a room
type CreateRoomRequest struct {
	Username string `json:"username"`
	Rounds   int    `json:"rounds"`
}

func generateRoomCode() string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}


func CreateRoom(c *fiber.Ctx) error {
	
	// Parse request body (JSON) (Needs username and rounds(Optional))
	var req CreateRoomRequest
	if err := c.BodyParser(&req); err != nil || req.Username == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Generate session ID
	sessionID := uuid.NewString()

	// Create player (owner)
	player := models.Player{
		Username:  req.Username,
		SessionID: sessionID,
	}
	if err := database.DB.Create(&player).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create player"})
	}

	// Generate unique room code
	roomCode := generateRoomCode()

	// Create room
	room := models.Room{
		Code:     roomCode,
		OwnerID:  player.ID,
		Rounds:   req.Rounds,
		// Status: default status is waiting 
	}
	if err := database.DB.Create(&room).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create room"})
	}

	// Add player to room (room owner)
	roomPlayer := models.RoomPlayer{
		RoomID:    room.ID,
		PlayerID:  &player.ID,
		Username:  player.Username,
		SessionID: sessionID,
	}
	if err := database.DB.Create(&roomPlayer).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to add owner to room"})
	}

	// Return room code and session ID
	return c.JSON(fiber.Map{
		"message":    "Room created successfully",
		"room_code":  roomCode,
		"session_id": sessionID,
	})
}

func JoinRoom(c *fiber.Ctx) error {
	// TODO: Join room logic
	return c.JSON(fiber.Map{"message": "Joined room"})
}

type StartGameRequest struct {
	RoomCode string `json:"room_code"`
}

func StartGame(c *fiber.Ctx) error {
	// Parse request body (JSON) (Needs room code)
	var req StartGameRequest
	if err := c.BodyParser(&req); err != nil || req.RoomCode == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Find room by code 
	var room models.Room
	if err := database.DB.Where("code = ?", req.RoomCode).First(&room).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Room not found"})
	}

	// Check if already started
	if room.Status == "in_progress" {
		return c.Status(400).JSON(fiber.Map{"error": "Game already started"})
	}

	// Get players in room
	var players []models.RoomPlayer
	if err := database.DB.Where("room_id = ?", room.ID).Find(&players).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch players"})
	}

	// Check if enough players
	if len(players) < 2 {
		return c.Status(400).JSON(fiber.Map{"error": "Need at least 2 players to start"})
	}

	// Make a player AI
	rand.Seed(time.Now().UnixNano())
	aiIndex := rand.Intn(len(players))
	players[aiIndex].IsAI = true
	database.DB.Save(&players[aiIndex])

	// Update room status
	room.Status = "in_progress"
	database.DB.Save(&room)

	// TODO: Broadcast to WebSocket room that game started
	
	// Return response (AI ID and message)
	return c.JSON(fiber.Map{
		"message": "Game started",
		"ai_id":   players[aiIndex].ID, // You can remove this from response later
	})
}