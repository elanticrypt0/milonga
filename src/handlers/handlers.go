package handlers

import (
	"fmt"
	"milonga/pkg/app"
	"milonga/pkg/utils"
	"net/http"
	"time"

	"github.com/elanticrypt0/dbman"
	"github.com/labstack/echo/v4"
)

func Setup(app *app.App) {

	/* DB Instance */
	// remove comments to use
	db := dbman.New()
	dbConfigPath := utils.GetAppRootPath() + "/config/db_config.toml"
	fmt.Printf("DB config path: %q\n", dbConfigPath)
	db.SetRootPath("./.db")
	db.LoadConfigToml(dbConfigPath)
	db.Connect("local")
	db.SetPrimary("local")

	app.DB = db

	// static path
	app.Server.Static("/", "./public")

	// para agrupar las urls en api
	api := app.Server.Group("api")

	api.GET("/example-htmlx", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<p>Hi from htmlx</p>")
	})

	api.GET("/example-clock", func(c echo.Context) error {

		hours, minutes, seconds := time.Now().Clock()
		currentTime := fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
		return c.HTML(http.StatusOK, currentTime)
	})

	api.GET("/example-load-clock", func(c echo.Context) error {

		return c.HTML(http.StatusOK, "<div hx-get=\"/api/example-clock\" hx-trigger=\"every 1s\">00:00:00</div>")
	})

	api.POST("/example-sayhi", func(c echo.Context) error {
		name := c.FormValue("name")
		return c.String(http.StatusOK, "Hi, "+name)
	})

}
