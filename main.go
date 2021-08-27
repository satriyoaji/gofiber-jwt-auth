package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gofiber-jwt-auth/database"
	"gofiber-jwt-auth/routes"
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