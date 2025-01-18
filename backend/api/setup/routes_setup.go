package setup

import (
	"milonga/api/routes"
	"milonga/internal/app"
)

func RoutesSetup(app *app.App) {

	router := app.Server.Group("api/v1")
	adminRouter := app.Server.Group("admin")

	// SetupMilongaRoutes(app, router)

	// routes.ProtectedRoutes(app)
	// routes.ExamplesRoutes(app)
	routes.AuthRoutes(app, adminRouter)
	routes.AdminRoutes(app, adminRouter)
	SetupMilongaRoutes(app, router)
	staticRoutes(app)

}

func staticRoutes(app *app.App) {
	app.Server.Static("/", "./web")
}
