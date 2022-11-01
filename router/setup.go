package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/soolllv/densys_milestone2/controllers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/register", controllers.Register)
}
