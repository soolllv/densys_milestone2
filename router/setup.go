package router

import (
	"github.com/gofiber/fiber/v2"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("SUP MATE")
}

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", hello)
}
