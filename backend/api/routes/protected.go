package routes

import (
	"milonga/api/handlers"
	"milonga/milonga/app"
	"milonga/milonga/vigilante"
)

func ProtectedRoutes(app *app.App) {

	middleware := vigilante.NewVigilanteMiddelware(app)
	protectedHandler := handlers.NewProtectedHander(app, app.DB.Primary)
	routes := app.Server.Group("/protected", middleware.IsLogged())

	// routes
	routes.Get("/", middleware.IsLogged(), protectedHandler.Index)
}
