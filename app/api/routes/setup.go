package routes

import (
	"milonga/internal/app"
	"milonga/pkg/vigilante"
)

func RoutesSetup(app *app.App) {

	vigilante.ActivateRoutes(app)

	protectedRoutes(app)
	examplesRoutes(app)
	staticRoutes(app)

}

func staticRoutes(app *app.App) {
	app.Server.Static("/", "./web")
}
