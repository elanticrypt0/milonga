package handlers

import (
	"milonga/internal/app"
	"milonga/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProtectedHander struct {
	handlers.Base
}

// NewUsersHandler crea una nueva instancia de UsersHandler
func NewProtectedHander(app *app.App, DB *gorm.DB) *ProtectedHander {
	return &ProtectedHander{
		Base: handlers.NewBaseHandler(app, DB),
	}
}

// handler de ejemplo
func (me *ProtectedHander) Index(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Welcome to protected route",
	})
}
