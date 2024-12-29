package routes

import (
	"milonga/api/handlers"
	"milonga/internal/app"
	"milonga/pkg/vigilante"

	"github.com/gofiber/fiber/v2"
)

func protectedRoutes(app *app.App) {

	middleware := vigilante.NewVigilanteMiddelware(app)

	protected := app.Server.Group("/protected", middleware.IsLogged())
	protected.Get("/index", func(c *fiber.Ctx) error {
		return handlers.ProtectedIndex(c, app)
	})
}
