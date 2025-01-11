package cli

import (
	"milonga/internal/app"
	"milonga/internal/vigilante"

	"gorm.io/gorm"
)

func VigilanteMigrate(app *app.App, db *gorm.DB) {
	db.AutoMigrate(&vigilante.User{})
}

func VigilanteAddAdmin(app *app.App, db *gorm.DB) {
	vigilante.CreateDefaultAdmin(db, app)
}

func VigilanteAddDefaultGuest(app *app.App, db *gorm.DB) {
	vigilante.CreateDefaultGuest(db, app)
}
