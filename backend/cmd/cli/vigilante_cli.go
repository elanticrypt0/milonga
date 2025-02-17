package cli

import (
	"log"
	"milonga/milonga/app"
	"milonga/milonga/vigilante"

	"gorm.io/gorm"
)

func VigilanteMigrate(db *gorm.DB) {
	log.Println("Installing vigilante basic...")
	db.AutoMigrate(&vigilante.UserAuth{})
	log.Println("All done!")
}

func VigilanteMigrateFull(db *gorm.DB) {
	log.Println("Installing vigilante...")

	VigilanteMigrate(db)
	VigilanteMigrateLoginAudit(db)
	VigilanteMigratePasswordToken(db)

	log.Println("All done!")
}

func VigilanteMigratePasswordToken(db *gorm.DB) {
	db.AutoMigrate(&vigilante.PasswordToken{})
	log.Println("Installing vigilante password token... migrated")
}

func VigilanteMigrateLoginAudit(db *gorm.DB) {
	db.AutoMigrate(&vigilante.LoginAudit{})
	log.Println("Installing vigilante audit... migrated")
}

func VigilanteAddAdmin(app *app.App, db *gorm.DB) {
	vigilante.CreateDefaultAdmin(db, app)
}

func VigilanteAddDefaultGuest(app *app.App, db *gorm.DB) {
	vigilante.CreateDefaultGuest(db, app)
}
