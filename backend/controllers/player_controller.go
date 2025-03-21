package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func ReconnectPlayer(c *fiber.Ctx) error {
	// TODO: Lookup player by session_id and rejoin room
	return c.JSON(fiber.Map{"message": "Player reconnected"})
}