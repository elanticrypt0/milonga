package vigilante

import (
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"milonga/internal/app"
)

type User struct {
	ID       uuid.UUID  `gorm:"type:uuid;primary_key;"`
	Email    string     `gorm:"unique;not null"`
	Username string     `gorm:"unique;not null"`
	Password string     `gorm:"type:text;not null"`
	Role     UserRole   `gorm:"type:varchar(20);default:'user'"`
	Status   UserStatus `gorm:"type:varchar(20);default:'active'"`
	gorm.Model
}

// BeforeCreate será llamado por GORM antes de crear un nuevo usuario
func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}

func (u *User) GetProfile(db *gorm.DB, userID string) error {
	result := db.First(u, "id = ?", userID)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
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
