package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {

	return c.String(http.StatusOK, "Hello, World!")
}

func HtmlxExample(c echo.Context) error {
	return c.HTML(http.StatusOK, "<p>Hi from htmlx</p>")
}

func SayHiExample(c echo.Context) error {
	name := c.FormValue("name")
	return c.String(http.StatusOK, "Hi, "+name)
}
