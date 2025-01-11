package cli

import (
	"fmt"
	"log"
	"milonga/database/migrations"
	"milonga/database/seeds"
	"milonga/internal/app"
	"milonga/internal/vigilante"

	"gorm.io/gorm"
)

func Setup() {
	PrintBanner("Milonga CLI", "0.1.46")
}

func Migrate(app *app.App, db *gorm.DB) {
	// App migrations
	migrations.AutoMigrate(db)
}

func Seed(db *gorm.DB) {
	fmt.Println("Seeding database")
	seeds.SeedCountries(db)
	fmt.Println("Seeding database... done")
}

func GenerateEncryptionKey(app *app.App) {
	key, err := vigilante.GenerateEncryptionKey()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Your new encryption key: %s\n", key)
	fmt.Println("Changer you app_config.toml file with this new key!")
}
