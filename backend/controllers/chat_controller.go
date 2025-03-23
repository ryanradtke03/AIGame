package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func SendMessage(c *fiber.Ctx) error {
	// TODO: Store message and broadcast to room
	return c.JSON(fiber.Map{"message": "Message sent"})
}

func GetChatHistory(c *fiber.Ctx) error {
	// TODO: Fetch chat history for current round
	return c.JSON(fiber.Map{"messages": []string{}})
}