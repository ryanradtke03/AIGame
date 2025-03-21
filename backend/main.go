package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"backend/routes"
	"backend/database"
)

func main() {

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// Initialize Fiber app
	app := fiber.New()

	// Middleware
	app.Use(logger.New()) // Logger Middleware
	app.Use(cors.New()) // CORS Middleware

	// // Connect to database
	database.ConnectDB()

	// Auto-migrate tables
	database.AutoMigrateTables()

	// Test Route (Fix: Ensure a valid root `/` route)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🚀 Welcome to Go + Fiber + PostgreSQL API!")
	})

	// Routes
	routes.RegisterRoutes(app)

	// Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))

}