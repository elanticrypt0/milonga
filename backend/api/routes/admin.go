package routes

import (
	"milonga/api/handlers"
	"milonga/milonga/app"
	"milonga/milonga/vigilante"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *app.App, router fiber.Router) {

	// middleware := vigilante.NewVigilanteMiddelware(app)
	authHandler := handlers.NewAuthHandler(app, app.DB.Primary)
	routes := router.Group("/auth")

	// routes
	routes.Get("/", authHandler.Index)
	routes.Post("/check", authHandler.Check)
	routes.Get("/logout", authHandler.Logout)
}

func AdminRoutes(app *app.App, router fiber.Router) {

	middleware := vigilante.NewVigilanteMiddelware(app)
	adminHandler := handlers.NewAdminHandler(app, app.DB.Primary)
	routes := router.Group("/panel", middleware.IsLogged(), middleware.IsStaff())

	// routes
	routes.Get("/", adminHandler.Index)
}
