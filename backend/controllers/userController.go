package controllers

import (
	"backend/database"
	"backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT id, name FROM users")
	if err != nil {
		return c.Status(500).SendString("Database error")
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return c.Status(500).SendString("Error scanning row")
		}
		users = append(users, user)
	}

	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	_, err := database.DB.Exec("INSERT INTO users (name) VALUES ($1)", user.Name)
	if err != nil {
		return c.Status(500).SendString("Error inserting user")
	}

	return c.Status(201).JSON(user)
}