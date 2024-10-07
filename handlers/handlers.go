package handlers

import (
	"milonga/db"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {

	db.ConnectDB()
	return c.SendString("Hello, World!")
}
