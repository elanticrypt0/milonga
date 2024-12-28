package src

import (
	"milonga/pkg/app"
)

func AppSetup(app *app.App) {

	/* REMOVE COMMENTS BELLOW TO USE DB */

	app.UseDB()
	app.DB.Connect("local")
	app.DB.SetPrimary("local")

}
