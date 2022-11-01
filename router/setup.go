package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/soolllv/densys_milestone2/controllers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/register-user", controllers.RegisterUser)
	api.Post("/register-doctor", controllers.RegisterDoctor)
}
