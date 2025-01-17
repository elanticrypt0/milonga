package setup

import (
	"milonga/api/routes"
	"milonga/internal/app"
)

func RoutesSetup(app *app.App) {

	router := app.Server.Group("api/v1")

	SetupMilongaRoutes(app, router)

	routes.ProtectedRoutes(app)
	routes.ExamplesRoutes(app)
	staticRoutes(app)

}

func staticRoutes(app *app.App) {
	app.Server.Static("/", "./web")
}
