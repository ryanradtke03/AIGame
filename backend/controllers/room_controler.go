package controllers

import (
	"backend/database"
	"backend/models"
	"backend/ws"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

// Create Room -------------------------------------------------------------------------
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

// Join Room -------------------------------------------------------------------------
type JoinRoomRequest struct {
	Username string `json:"username"`
	Code     string `json:"code"`
}

func JoinRoom(c *fiber.Ctx) error {

	// Parse request body (JSON) (Needs username and room code)
	var req JoinRoomRequest
	if err := c.BodyParser(&req); err != nil || req.Username == "" || req.Code == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Find room by code
	var room models.Room
	if err := database.DB.Where("code = ?", req.Code).First(&room).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Room not found"})
	}

	// Check if room is full (max 6 players (for now))
	var count int64
	database.DB.Model(&models.RoomPlayer{}).Where("room+id = ?", room.ID).Count(&count)
	if count >= 6 {
		return c.Status(400).JSON(fiber.Map{"error": "Room is full"})
	}

	// Create a new player
	sessionID := uuid.NewString()
	player := models.Player{
		Username:  req.Username,
		SessionID: sessionID,
	}
	if err := database.DB.Create(&player).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create player"})
	}

	// Add player to room
	roomPlayer := models.RoomPlayer{
		RoomID:    room.ID,
		PlayerID:  &player.ID,
		Username:  player.Username,
		SessionID: sessionID,
	}
	if err := database.DB.Create(&roomPlayer).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to add player to room"})
	}

	// Broadcast "player_joined" to everyone in the room
	ws.BroadcastJSON(room.Code, map[string]string{
		"event":    "player_joined",
		"username": req.Username,
	})

	// Return response
	return c.JSON(fiber.Map{
		"message":    "Joined room successfully",
		"room_code":  room.Code,
		"session_id": sessionID,
	})
}

// Start Room / Game -------------------------------------------------------------------------
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

	// Create initial game round (Round 1)
	gameRound := models.GameRound{
		RoomID:      room.ID,
		RoundNumber: 1,
		AIRevealed:  false,
	}
	if err := database.DB.Create(&gameRound).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create initial round"})
	}

	// Broadcast "game_started" to everyone in the room
	ws.BroadcastJSON(room.Code, map[string]string{
		"event":   "game_started",
		"message": "The game has begun!",
	})

	// Start round timer (60 seconds for now) (Can be changed later) (Maybe a setting)
	go startRoundTimer(room.Code, gameRound.ID, 60*time.Second)
	
	// Return response (AI ID and message)
	return c.JSON(fiber.Map{
		"message": "Game started",
		"ai_id":   players[aiIndex].ID, // You can remove this from response later
	})
}


func startRoundTimer(roomCode string, roundID uint, duration time.Duration) {
	time.Sleep(duration)

	// TODO: mark round as ended (set AIRevealed true later)
	var round models.GameRound
	if err := database.DB.First(&round, roundID).Error; err == nil {
		// TODO: Logic to update the round phase here if needed
	}

	// Broadcast voting started
	ws.BroadcastJSON(roomCode, map[string]string{
		"event":   "voting_started",
		"message": "Time's up! Voting has begun.",
	})
}