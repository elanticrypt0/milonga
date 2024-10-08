package app

import (
	"github.com/elanticrypt0/dbman"
	"github.com/labstack/echo/v4"
)

type App struct {
	Server *echo.Echo
	Config *Config
	DB     *dbman.DBMan
}
