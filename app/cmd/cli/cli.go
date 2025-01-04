package cli

import (
	"fmt"
	"milonga/database/migrations"
	"milonga/internal/app"
	"milonga/pkg/vigilante"

	"gorm.io/gorm"
)

func Setup() {
	PrintBanner("Milonga CLI", "0.1.46")
}

func VigilanteMigrate(app *app.App, db *gorm.DB) {
	db.AutoMigrate(&vigilante.User{})
}

func VigilanteAddAdmin(app *app.App, db *gorm.DB) {
	vigilante.CreateDefaultAdmin(db, app)
}

func Migrate(app *app.App, db *gorm.DB) {
	// App migrations
	migrations.AutoMigrate(db)
}

func Seed(db *gorm.DB) {
	fmt.Println("Seeding database")
	fmt.Println("Seeding database... done")
}
