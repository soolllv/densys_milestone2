package controllers

import "github.com/gofiber/fiber/v2"

func hello(c *fiber.Ctx) error {
	return c.SendString("SUP MATE")
}
