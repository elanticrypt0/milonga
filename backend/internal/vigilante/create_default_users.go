package vigilante

import (
	"log"
	"milonga/internal/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateDefaultAdmin(db *gorm.DB, app *app.App) error {
	var count int64
	db.Model(&User{}).Where("role = ?", "admin").Count(&count)

	// Si ya existe al menos un admin, no creamos uno nuevo
	if count > 0 {
		log.Println("Admin user already exists")
		return nil
	}

	default_admin := app.LoadDefaultAdminConfig()

	// Obtener credenciales del admin desde variables de entorno
	adminEmail := default_admin.Email
	adminPassword := default_admin.Password
	adminUsername := default_admin.Username

	// Hash de la contraseña
	hashedPassword, err := HashPassword(adminPassword)
	if err != nil {
		return err
	}

	// Crear usuario admin
	admin := User{
		Email:    adminEmail,
		Username: adminUsername,
		Password: string(hashedPassword),
		Role:     "admin",
		Status:   UserStatusEnabled,
	}

	result := db.Create(&admin)
	if result.Error != nil {
		return result.Error
	}

	log.Printf("Default admin user created with email: %s\n", adminEmail)
	return nil
}

func CreateDefaultGuest(db *gorm.DB, app *app.App) error {
	// random username
	randomUsername := GenerateUsername(12, 24)

	var count int64
	db.Model(&User{}).Where("username = ?", randomUsername).Count(&count)

	// Obtener credenciales del admin desde variables de entorno
	vipGuestEmail := "guest@token.pass"
	vipGuestPassword := uuid.New().String()
	vipGuestUsername := randomUsername

	// Hash de la contraseña
	hashedPassword, err := HashPassword(vipGuestPassword)
	if err != nil {
		return err
	}

	// Crear usuario vipGuest
	vipGuest := User{
		Email:    vipGuestEmail,
		Username: vipGuestUsername,
		Password: string(hashedPassword),
		Role:     "user",
		Status:   UserStatusEnabled,
	}

	var vipGuest2 User

	result := db.Create(&vipGuest)
	if result.Error != nil {
		return result.Error
	}

	db.Find(&vipGuest2, "email=?", vipGuestEmail)

	// genera el token de acceso

	passtoken := NewPasswordToken()
	plaintoken, err := passtoken.Create(vipGuest2.ID, db)
	if err != nil {
		return err
	}

	log.Printf("Default guest user created with email: %s and password token : %q\n", vipGuestEmail, plaintoken)
	return nil
}
