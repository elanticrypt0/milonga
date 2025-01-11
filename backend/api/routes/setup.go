package routes

import (
	"milonga/internal/app"
)

func RoutesSetup(app *app.App) {

	router := app.Server.Group("api/v1")

	SetupMilongaRoutes(app, router)

	ProtectedRoutes(app)
	examplesRoutes(app)
	staticRoutes(app)

}

func staticRoutes(app *app.App) {
	app.Server.Static("/", "./web")
}
