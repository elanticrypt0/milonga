package handlers

import (
	"fmt"
	"milonga/pkg/app"
	"milonga/pkg/milonga_response"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *app.App) {

	// static path
	app.Server.Static("/", "./public")

	// para agrupar las urls en api
	api := app.Server.Group("api")

	api.Get("/htmx-to-astro", func(c *fiber.Ctx) error {

		return milonga_response.SendHTMLFromFile(c, "./public/examplex.html")

	})

	api.Get("/example-htmlx", func(c *fiber.Ctx) error {
		resp := "<p>Hi from htmlx</p>"
		// return c.Render("response-htmx", fiber.Map{
		// 	"Response": "",
		// })

		return milonga_response.SendHTML(c, resp)

	})

	api.Get("/hi-from-file", func(c *fiber.Ctx) error {

		return milonga_response.SendHTMLFromFile(c, app.Config.ViewsPath+"/hifromfile.html")

	})

	api.Get("/hi-from-file2", func(c *fiber.Ctx) error {

		return milonga_response.SendHTMLFromFile(c, app.Config.ViewsPath+"/hifromfile2.html")

	})

	api.Get("/example-clock", func(c *fiber.Ctx) error {

		hours, minutes, seconds := time.Now().Clock()
		currentTime := fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)

		return milonga_response.SendHTML(c, currentTime)
	})

	api.Get("/example-load-clock", func(c *fiber.Ctx) error {
		resp := "<div hx-get=\"/api/example-clock\" hx-trigger=\"every 1s\">00:00:00</div>"

		return milonga_response.SendHTML(c, resp)
	})

	api.Post("/example-sayhi", func(c *fiber.Ctx) error {
		name := c.FormValue("name")

		resp := "Hi, " + name

		return milonga_response.SendHTML(c, resp)
	})

}
