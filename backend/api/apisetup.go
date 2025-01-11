package src

import (
	"milonga/api/routes"
	"milonga/internal/app"
	"milonga/internal/vigilante"
)

func ApiSetup(app *app.App) {

	app.UseDB()
	app.DB.Connect("local")
	app.DB.SetPrimary("local")

	app.DB.Primary.AutoMigrate(&vigilante.User{})
	// remove in production
	vigilante.CreateDefaultAdmin(app.DB.Primary, app)

	routes.RoutesSetup(app)

}
