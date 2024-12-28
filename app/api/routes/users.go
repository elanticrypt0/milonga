package routes

import (
	"milonga/pkg/app"
	"milonga/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func usersRoutes(app *app.App){
	users := app.Server.Group("/users")
	users.Get("/", func(c *fiber.Ctx) error{
		return handlers.GetAllUsers(c,app)
	})
	users.Get("/:id", func(c *fiber.Ctx) error{
		return handlers.GetUser(c,app)
	})
	users.Get("/search", func(c *fiber.Ctx) error{
		return handlers.SearchUser(c,app)
	})
	users.Put("/:id", func(c *fiber.Ctx) error{
		return handlers.UpdateUser(c,app)
	})
	users.Delete("/:id", func(c *fiber.Ctx) error{
		return handlers.DeleteUser(c,app)
	})
}