// Copyright (c) 2024 Milonga
// This software is licensed under the MIT License.
// See LICENSE file in the project root for full license information.

package main

import (
	"context"
	"fmt"
	"github.com/gofiber/template/html/v2"
	"log"
	"milonga/api/setup"
	"milonga/milonga/app"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	// read env config
	// config := app.LoadConfig("./config/app_config.toml")

	app := app.New("./config/app_config.toml")
	app.SetCtx(context.Background())

	templateEngine := html.New(app.Config.TemplatesPath, ".html")

	// Create a new engineRender to render HTML

	app.Server = fiber.New(fiber.Config{
		Prefork:           false,
		CaseSensitive:     true,
		StrictRouting:     true,
		EnablePrintRoutes: false,
		ServerHeader:      app.Config.Name,
		AppName:           app.Config.Name + " v" + app.Config.Version,
		Views:             templateEngine,
	})

	allowedCORS := fmt.Sprintf("%s, %s:4321, %s", app.Config.AppHost, app.Config.URL, app.Config.AllowedCORS)
	fmt.Printf("ALLOWED CORS: %s\n", allowedCORS)

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

	setup.ApiSetup(app)

	app.Server.Use(recover.New())

	log.Fatal(app.Server.Listen(":" + app.Config.Port))
}
