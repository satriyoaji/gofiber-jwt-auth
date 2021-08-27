package main

import (
	"fiber-jwt-auth/database"
	"fiber-jwt-auth/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	})) // like different port between backend and frontend

	routes.Setup(app)

	app.Listen(":8000")
}