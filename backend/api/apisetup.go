package src

import (
	"milonga/api/routes"
	"milonga/internal/app"
)

func ApiSetup(app *app.App) {

	DatabaseSetup(app)
	routes.RoutesSetup(app)

}

// also used in CMD
func DatabaseSetup(app *app.App) {
	app.UseDB()
	// app.DB.Connect("local")
	// app.DB.SetPrimary("local")

	app.DB.Connect("mysql-test")
	app.DB.SetPrimary("mysql-test")
}
