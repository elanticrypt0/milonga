package routes

import (
	"fmt"
	"milonga/internal/app"
	"milonga/internal/milonga_response"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ExamplesRoutes(app *app.App) {

	// para agrupar las urls en api
	api := app.Server.Group("api")

	api.Get("/htmx-to-astro", func(c *fiber.Ctx) error {

		return milonga_response.SendHTMLFromFile(c, "./web/examplex.html")

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

	// ejemplo de utilizacion de plantillas repetidas
	api.Get("/example-several-records", func(c *fiber.Ctx) error {

		personas := []struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}{
			{Name: "John Doe", Email: "john@example.com"},
			{Name: "Jane Doe", Email: "jane@example.com"},
			{Name: "Jane Doe", Email: "jane@example.com"},
			{Name: "Jane Doe", Email: "jane@example.com"},
		}

		html, err := milonga_response.PrepareHTMLFromSlice("./web/views/personas.html", personas)
		if err != nil {
			return milonga_response.SendHTML(c, "No hay ninguna persona para mostrar")
		}
		return milonga_response.SendHTML(c, string(html))
	})

}
