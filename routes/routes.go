package routes

import (
	"github.com/gofiber/fiber/v2"
)

func IndexHandler(c *fiber.Ctx) error {
	return c.SendString("Hell Xoxoxoxo")
}

func ToolHandler(c *fiber.Ctx) error {
	return c.SendString("Tools Cococococ")
}
