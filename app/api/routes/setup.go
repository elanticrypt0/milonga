package routes

import (
	"milonga/pkg/app"
)

func RoutesSetup(app *app.App) {

	usersRoutes(app)
	authRoutes(app)
	protectedRoutes(app)
	examplesRoutes(app)
	staticRoutes(app)

}

func staticRoutes(app *app.App) {
	app.Server.Static("/", "./web")
}
