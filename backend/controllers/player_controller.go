package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"backend/database"
	"backend/models"
)

func ReconnectPlayer(c *fiber.Ctx) error {
	// TODO: Lookup player by session_id and rejoin room
	return c.JSON(fiber.Map{"message": "Player reconnected"})
}


type RoomPlayerDTO struct {
	Username string `json:"username"`
	IsAI     bool   `json:"isAI"`
	Points   int    `json:"points"`
	Score    int    `json:"score"`
	IsOwner  bool   `json:"isOwner"`
	SessionID string `json:"sessionID"`
}

func GetPlayersInRoom(c *fiber.Ctx) error {
	// Parse room ID from URL
	roomIDParam := c.Params("roomID")
	roomIDUint64, err := strconv.ParseUint(roomIDParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid room ID",
		})
	}
	roomID := uint(roomIDUint64)

	// Get the room to find the owner
	var room models.Room
	if err := database.DB.First(&room, roomID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Room not found",
		})
	}

	// Get all room players
	var roomPlayers []models.RoomPlayer
	if err := database.DB.
		Where("room_id = ?", roomID).
		Find(&roomPlayers).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch players",
		})
	}

	// Map to trimmed-down DTOs
	var response []RoomPlayerDTO
	for _, p := range roomPlayers {
		isOwner := p.PlayerID != nil && *p.PlayerID == room.OwnerID

		response = append(response, RoomPlayerDTO{
			Username: p.Username,
			IsAI:     p.IsAI,
			Points:   p.Points,
			Score:    p.Score,
			IsOwner:  isOwner,
			SessionID: p.SessionID,
		})
	}

	return c.JSON(response)
}
