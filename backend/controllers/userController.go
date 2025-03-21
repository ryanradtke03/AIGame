package controllers

import (
	"backend/database"
	"backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	// ✅ Correct way to fetch users in GORM
	result := database.DB.Raw("SELECT id, name FROM users").Scan(&users)
	if result.Error != nil {
		return c.Status(500).SendString("Database error")
	}

	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	// ✅ Correct way to insert users in GORM
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(500).SendString("Error inserting user")
	}

	return c.Status(201).JSON(user)
}