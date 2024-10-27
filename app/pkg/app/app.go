package app

import (
	"fmt"
	"milonga/pkg/utils"

	"github.com/elanticrypt0/dbman"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	Server *fiber.App
	Config *Config
	DB     *dbman.DBMan
}

func New(configPath string) *App {

	app := &App{
		Config: LoadConfig(configPath),
	}

	app.Config.AppHost = app.Config.URL + ":" + app.Config.Port

	// app.UseDB()

	return app
}

func (me *App) UseDB() {
	db := dbman.New()
	dbConfigPath := utils.GetAppRootPath() + me.Config.DBConfigPath
	fmt.Printf("DB config path: %q\n", me.Config.DBConfigPath)
	db.SetRootPath("./.db")
	db.LoadConfigToml(dbConfigPath)
	me.DB = db
}
