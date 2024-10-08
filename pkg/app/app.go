package app

import (
	"github.com/labstack/echo/v4"
)

type App struct {
	Server *echo.Echo
	Config *Config
}
