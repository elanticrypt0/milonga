package routes

import (
	"milonga/pkg/app"
	"milonga/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func authRoutes(app *app.App){
	auth := app.Server.Group("auth")

	auth.Post("/register", func (c *fiber.Ctx) error{
		return handlers.Register(c, app)
	})
    auth.Post("/login", func (c *fiber.Ctx) error{
		return handlers.Login(c,app)
	})

}