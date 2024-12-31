package healthcheck

import (
	"milonga/internal/app"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// explose GET /health and GET /ping
func ActivateRoutes(app *app.App, services []CriticalService, dbs2check []*gorm.DB) {

	healtchecker := NewHealthCheck(app, services, dbs2check)
	app.Server.Get("/health", healtchecker.Check)

	app.Server.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("pong")
	})

}
