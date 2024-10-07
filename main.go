package main

import (
	"milonga/app"
	"milonga/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// read env config

	config := app.LoadConfig("./config/app_config.toml")

	// inicia servicio
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       config.Name + " v" + config.Version,
	})

	app.Get("/", handlers.Index)

	app.Listen(":" + config.Port)
}
