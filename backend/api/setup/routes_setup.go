package setup

import (
	"milonga/api/routes"
	"milonga/milonga/app"
)

func RoutesSetup(app *app.App) {

	// default api path is api/v1
	router := app.Server.Group(app.Config.APIPath)
	// adminRouter := app.Server.Group("admin")

	// SetupMilongaRoutes(app, router)

	// routes.ProtectedRoutes(app)
	// routes.ExamplesRoutes(app)
	routes.TaskRoutes(app, router)

	SetupMilongaRoutes(app, router)

	staticRoutes(app)

}

func staticRoutes(app *app.App) {
	app.Server.Static("/", "./web")
}
