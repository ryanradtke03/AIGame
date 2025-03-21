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

func StartGame(c *fiber.Ctx) error {
	// TODO: Set room status to in_progress and assign AI
	return c.JSON(fiber.Map{"message": "Game started"})
}