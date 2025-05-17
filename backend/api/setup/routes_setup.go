package setup

import (
	"milonga/api/routes"
	"milonga/milonga/app"
)

func RoutesSetup(app *app.App) {

	router := app.Server.Group("api/v1")
	adminRouter := app.Server.Group("admin")

	SetupMilongaRoutes(app, router)

	//routes.ProtectedRoutes(app)
	//routes.ExamplesRoutes(app)

	//ENABLE TO use Authroutes on browser
	routes.AuthRoutes(app, adminRouter)

	//ENABLE TO use AdminRoutes on Browser
	routes.AdminRoutes(app, adminRouter)

	//SetupMilongaRoutes(app, router)
	HealthRoutes(app)

	//staticRoutes(app)

}

func staticRoutes(app *app.App) {
	app.Server.Static("/", "./web")
}
