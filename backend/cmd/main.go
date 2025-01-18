package main

import (
	"flag"
	"fmt"
	"milonga/api/setup"
	"milonga/cmd/cli"
	"milonga/internal/app"
	"milonga/internal/utils"
	"os"
)

func main() {

	version := "0.1.50"
	cli.Setup(version)
	var app *app.App

	// no conenct to DB
	noLoadApp := flag.Bool("nc", false, "No connect to database")
	// vigilante commands
	generateVigilanteEncryptionkeyCmd := flag.Bool("vigilante:generate_key", false, "Run database migrations")
	migrateVigilanteBasicCmd := flag.Bool("vigilante:install", false, "Run database migrations")
	migrateVigilanteFullCmd := flag.Bool("vigilante:install_full", false, "Run database migrations")
	migrateVigilanteGuestCmd := flag.Bool("vigilante:make_guest", false, "Run database migrations")
	// migrations & seeds
	migrateCmd := flag.Bool("migrate", false, "Run database migrations")
	seedCmd := flag.Bool("seed", false, "Run database seeds")

	flag.Parse()

	if !*noLoadApp {
		app = appSetup()
	}

	if !*migrateCmd && !*migrateVigilanteBasicCmd && !*migrateVigilanteFullCmd && !*migrateVigilanteGuestCmd && !*generateVigilanteEncryptionkeyCmd && !*seedCmd {
		fmt.Println("Uso: programa -migrate o -seed")
		os.Exit(1)
	}

	if *migrateVigilanteBasicCmd {
		cli.VigilanteMigrateFull(app.DB.Primary)
		cli.VigilanteAddAdmin(app, app.DB.Primary)
	}

	if *migrateVigilanteFullCmd {
		cli.VigilanteMigrateFull(app.DB.Primary)
		cli.VigilanteAddAdmin(app, app.DB.Primary)
		cli.VigilanteAddDefaultGuest(app, app.DB.Primary)
	}

	if *migrateVigilanteGuestCmd {
		cli.VigilanteAddDefaultGuest(app, app.DB.Primary)
	}

	if *generateVigilanteEncryptionkeyCmd {
		cli.GenerateEncryptionKey(app)
	}

	if *migrateCmd {
		cli.Migrate(app, app.DB.Primary)
	}

	if *seedCmd {
		cli.Seed(app.DB.Primary)
	}
}

func appSetup() *app.App {
	app := app.New(utils.GetAppRootPath() + "/config/app_config.toml")

	// call default api databaseconnection
	setup.DatabaseSetup(app)

	return app
}
