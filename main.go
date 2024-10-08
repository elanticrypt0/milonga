package main

import (
	"milonga/handlers"
	"milonga/pkg/app"
	"milonga/pkg/utils"

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
		AllowOrigins: []string{"http://localhost:" + app.Config.Port, "http://localhost:4321", "http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	app.Server.Use(middleware.Logger())
	app.Server.Use(middleware.Recover())

	handlers.Setup(app)

	utils.OpenInBrowser("http://localhost:" + app.Config.Port)

	app.Server.Logger.Fatal(app.Server.Start(":" + app.Config.Port))
}
