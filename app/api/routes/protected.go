package routes

import (
	"milonga/api/handlers"
	"milonga/pkg/app"
	"milonga/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func protectedRoutes(app *app.App) {
	protected := app.Server.Group("/protected", middleware.IsLogged(app))
	protected.Get("/index", func(c *fiber.Ctx) error {
		return handlers.ProtectedIndex(c, app)
	})
}
