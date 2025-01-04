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
	migrateVigilanteCmd := flag.Bool("migrate-vigilante", false, "Run database migrations")
	seedCmd := flag.Bool("seed", false, "Run database seeds")

	flag.Parse()

	if !*migrateCmd && !*migrateVigilanteCmd && !*seedCmd {
		fmt.Println("Uso: programa -migrate o -seed")
		os.Exit(1)
	}

	if *migrateVigilanteCmd {
		cli.MigrateVigilante(app, app.DB.Primary)
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
