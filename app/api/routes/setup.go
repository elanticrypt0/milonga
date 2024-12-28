package routes

import(
	"milonga/pkg/app"
)

func RoutesSetup(app *app.App){
	
	app.Server.Static("/", "./web")
	
	// api := app.Server.Group("/api/", middleware.Protected())

	usersRoutes(app)
	authRoutes(app)
	protectedRoutes(app)
	examplesRoutes(app)

}

func staticRoutes(app *app.App){
	app.Server.Static("/", "./web")
}
