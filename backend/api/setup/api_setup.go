package setup

import (
	"milonga/milonga/app"
	"milonga/milonga/utils"
)

func ApiSetup(app *app.App) {

	DatabaseSetup(app)
	RoutesSetup(app)

	if app.Config.OpenInBrowser {
		utils.OpenInBrowser(app.Config.AppHost)
	}

	app.ConsoleMessage("Running on " + app.Config.Enviroment)
}

// also used in CMD
func DatabaseSetup(app *app.App) {
	app.UseDB()
	app.DB.Connect("local")
	app.DB.SetPrimary("local")

	// app.DB.Connect("mysql_dev")
	// app.DB.SetPrimary("mysql_dev")

	// app.DB.Connect("psg_dev")
	// app.DB.SetPrimary("psg_dev")
}
