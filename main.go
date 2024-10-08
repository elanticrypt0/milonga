package main

import (
	"milonga/pkg/app"
	"milonga/pkg/utils"
	"milonga/src/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// read env config
	// config := app.LoadConfig("./config/app_config.toml")

	app := &app.App{
		Server: echo.New(),
		Config: app.LoadConfig("./config/app_config.toml"),
	}

	app.Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{app.Config.URL + ":" + app.Config.Port, "http://localhost:4321"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	app.Server.Use(middleware.Logger())
	app.Server.Use(middleware.Recover())

	handlers.Setup(app)

	utils.OpenInBrowser(app.Config.URL + ":" + app.Config.Port)

	app.Server.Logger.Fatal(app.Server.Start(":" + app.Config.Port))
}
