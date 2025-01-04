package cli

import (
	"fmt"
	"milonga/internal/app"
)

func Setup() {
	PrintBanner("Milonga CLI", "0.1.46")
}

func Migrate(app *app.App) {
	fmt.Println("Running database migration")
}

func Seed(app *app.App) {
	fmt.Println("Seeding the database")
}
