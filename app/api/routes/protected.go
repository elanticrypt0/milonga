package routes

import(
	"milonga/pkg/app"
	"milonga/api/handlers"
	"milonga/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func protectedRoutes(app *app.App){
    protected := app.Server.Group("/protected", middleware.Protected(app))
    protected.Get("/index", func(c *fiber.Ctx) error {
        return handlers.ProtectedIndex(c, app)
    })
}