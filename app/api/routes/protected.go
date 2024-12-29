package routes

import (
	"milonga/api/handlers"
	"milonga/internal/app"
	"milonga/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func protectedRoutes(app *app.App) {

	middleware := middleware.NewUserAuthMiddelware(app)

	protected := app.Server.Group("/protected", middleware.IsLogged())
	protected.Get("/index", func(c *fiber.Ctx) error {
		return handlers.ProtectedIndex(c, app)
	})
}
