package routes

import (
	"milonga/api/handlers"
	"milonga/pkg/app"
	"milonga/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func usersRoutes(app *app.App) {
	users := app.Server.Group("/users", middleware.IsLogged(app), middleware.NotUser(app))

	users.Get("/", func(c *fiber.Ctx) error {
		return handlers.GetAllUsers(c, app)
	})

	users.Post("/", func(c *fiber.Ctx) error {
		return handlers.CreateUser(c, app)
	})

	users.Get("/profile", func(c *fiber.Ctx) error {
		return handlers.GetProfile(c, app)
	})

	users.Get("/:id", func(c *fiber.Ctx) error {
		return handlers.GetUser(c, app)
	})
	users.Get("/search", func(c *fiber.Ctx) error {
		return handlers.SearchUser(c, app)
	})
	users.Put("/:id", func(c *fiber.Ctx) error {
		return handlers.UpdateUser(c, app)
	})
	users.Delete("/:id", middleware.RequireRole(app, "admin"), func(c *fiber.Ctx) error {
		return handlers.DeleteUser(c, app)
	})
}
