package main

import (
	"fmt"
	"log"
	src "milonga/api"
	"milonga/internal/app"
	"milonga/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	app := app.New(utils.GetAppRootPath() + "/config/app_config.toml")

	app.Server = fiber.New(fiber.Config{
		Prefork:           false,
		CaseSensitive:     true,
		StrictRouting:     true,
		EnablePrintRoutes: false,
		ServerHeader:      app.Config.Name,
		AppName:           app.Config.Name + " v" + app.Config.Version,
	})

	allowedCORS := fmt.Sprintf("%s, %s:4321", app.Config.AppHost, app.Config.URL)

	app.Server.Use(cors.New(cors.Config{
		AllowOrigins: allowedCORS,
		// AllowHeaders: "Origin, Content-Type, Accept",
	}))

	/* REMOVE COMMENTS BELOW TO USE A LOG FILE */
	// set log file's path

	// logFile, err := os.OpenFile(app.Config.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("error opening file: %v", err)
	// }
	// defer logFile.Close()

	app.Server.Use(logger.New(logger.Config{
		// Output:     logFile,
		Format:     "PID: ${pid} [${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/Argentina/Buenos_Aires",
	}))

	src.ApiSetup(app)

	// remove this to open the web on the start
	// utils.OpenInBrowser(app.Config.AppHost)

	app.Server.Use(recover.New())

	log.Fatal(app.Server.Listen(":" + app.Config.Port))
}
