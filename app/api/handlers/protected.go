package handlers

import (
	"milonga/internal/app"

	"github.com/gofiber/fiber/v2"
)

func ProtectedIndex(c *fiber.Ctx, app *app.App) error {
	return c.JSON(fiber.Map{
		"message": "Welcome to protected route",
	})
}
