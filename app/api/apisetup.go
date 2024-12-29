package src

import (
	"milonga/api/models"
	"milonga/api/routes"
	"milonga/internal/app"
)

func ApiSetup(app *app.App) {

	app.UseDB()
	app.DB.Connect("local")
	app.DB.SetPrimary("local")

	app.DB.Primary.AutoMigrate(&models.User{})
	// remove in production
	models.CreateDefaultAdmin(app.DB.Primary, app)

	routes.RoutesSetup(app)

}
