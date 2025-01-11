package routes

import (
	"milonga/internal/app"
)

func RoutesSetup(app *app.App) {

	SetupMilongaRoutes(app)

	protectedRoutes(app)
	examplesRoutes(app)
	staticRoutes(app)

}

func staticRoutes(app *app.App) {
	app.Server.Static("/", "./web")
}
