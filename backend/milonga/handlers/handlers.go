package handlers

import (
	"milonga/milonga/app"

	"gorm.io/gorm"
)

type Base struct {
	App *app.App
	DB  *gorm.DB
}

func NewBaseHandler(app *app.App, DB *gorm.DB) Base {
	return Base{
		App: app,
		DB:  DB,
	}
}
