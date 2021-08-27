package routes

import (
	"github.com/gofiber/fiber/v2"
	"gofiber-jwt-auth/controllers"
)

func Setup(app *fiber.App)  {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/users", controllers.User)
	app.Post("/api/logout", controllers.Logout)
}
