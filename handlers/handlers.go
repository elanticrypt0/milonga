package handlers

import (
	"milonga/pkg/app"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Setup(app *app.App) {

	/* DB Instance */
	// remove comments to use
	// db := db.ConnectDB("local")
	// db.SetPrimary("local")
	// db.Primary.AutoMigrate(&models.File{})

	// static path
	app.Server.Static("/public", "./public")

	app.Server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// para agrupar las urls en api
	api := app.Server.Group("api")

	api.GET("/example-htmlx", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<p>Hi from htmlx</p>")
	})

	api.POST("/example-sayhi", func(c echo.Context) error {
		name := c.FormValue("name")
		return c.String(http.StatusOK, "Hi, "+name)
	})

}
