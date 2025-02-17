package app

import (
	"fmt"
	"log"
	"milonga/milonga/utils"

	"milonga/milonga/dbman"

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

	return app
}

func (me *App) UseDB() {
	db := dbman.New()
	dbConfigPath := utils.GetAppRootPath() + me.Config.DBConfigPath
	fmt.Printf("DB config path: %q\n", me.Config.DBConfigPath)
	db.SetRootPath(utils.GetAppRootPath() + "/database")
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

func (me *App) GetCurrentEnviroment() string {
	return me.Config.Enviroment
}

func (me *App) ConsoleMessage(message string) {
	log.Printf("Milonga :: %s\n", message)
}
