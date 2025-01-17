package app

import (
	"fmt"
	"milonga/internal/utils"

	"milonga/internal/dbman"

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
	db.SetRootPath("./database")
	db.LoadConfigToml(dbConfigPath)
	me.DB = db
}

// load default userData
func (me *App) LoadDefaultAdminConfig() *DefaultAdmin {
	admin := &DefaultAdmin{}
	daConfigPath := utils.GetAppRootPath() + "/config/default_admin.toml"
	LoadTomlFile(daConfigPath, admin)
	return admin
}
