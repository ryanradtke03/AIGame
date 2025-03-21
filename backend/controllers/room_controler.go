package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func CreateRoom(c *fiber.Ctx) error {
	// TODO: Create room logic
	return c.JSON(fiber.Map{"message": "Room created"})
}

func JoinRoom(c *fiber.Ctx) error {
	// TODO: Join room logic
	return c.JSON(fiber.Map{"message": "Joined room"})
}

func StartGame(c *fiber.Ctx) error {
	// TODO: Set room status to in_progress and assign AI
	return c.JSON(fiber.Map{"message": "Game started"})
}