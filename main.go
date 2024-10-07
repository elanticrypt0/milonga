package main

import (
	"milonga/app"
	"milonga/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// read env config
	config := app.LoadConfig("./config/app_config.toml")

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:" + config.Port, "http://localhost:4321", "http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	/* Handlers */
	e.Static("/public", "./public")

	e.GET("/", handlers.Index)

	e.GET("/example-htmlx", handlers.HtmlxExample)

	e.Logger.Fatal(e.Start(":" + config.Port))
}
