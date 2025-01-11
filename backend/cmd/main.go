package main

import (
	"flag"
	"fmt"
	"milonga/cmd/cli"
	"milonga/internal/app"
	"milonga/internal/utils"
	"os"
)

func main() {

	cli.Setup()
	app := appSetup()

	migrateCmd := flag.Bool("migrate", false, "Run database migrations")
	migrateVigilanteBasicCmd := flag.Bool("vigilante:install_basic", false, "Run database migrations")
	migrateVigilanteFullCmd := flag.Bool("vigilante:install_full", false, "Run database migrations")
	migrateVigilanteGuestCmd := flag.Bool("vigilante:make_guest", false, "Run database migrations")
	generateVigilanteEncryptionkeyCmd := flag.Bool("vigilante:generate_key", false, "Run database migrations")
	seedCmd := flag.Bool("seed", false, "Run database seeds")

	flag.Parse()

	if !*migrateCmd && !*migrateVigilanteBasicCmd && !*migrateVigilanteFullCmd && !*migrateVigilanteGuestCmd && !*generateVigilanteEncryptionkeyCmd && !*seedCmd {
		fmt.Println("Uso: programa -migrate o -seed")
		os.Exit(1)
	}

	if *migrateVigilanteBasicCmd {
		cli.VigilanteMigrate(app.DB.Primary)
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

	/* Change this with your needs */
	app.UseDB()
	app.DB.Connect("local")
	app.DB.SetPrimary("local")

	return app
}
