package models

import (

	"golang.org/x/crypto/bcrypt"
    "github.com/google/uuid"
    "gorm.io/gorm"
	"log"

    "milonga/pkg/app"
)

type User struct {
    ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
    Email     string    `gorm:"unique;not null"`
    Username  string    `gorm:"unique;not null"`
    Password  string    `gorm:"not null"`
    Role      string    `gorm:"type:varchar(20);default:'user'"`
    gorm.Model
}

// BeforeCreate será llamado por GORM antes de crear un nuevo usuario
func (u *User) BeforeCreate(tx *gorm.DB) error {
    u.ID = uuid.New()
    return nil
}

func CreateDefaultAdmin(db *gorm.DB,app *app.App) error {
    var count int64
    db.Model(&User{}).Where("role = ?", "admin").Count(&count)
    
    // Si ya existe al menos un admin, no creamos uno nuevo
    if count > 0 {
        log.Println("Admin user already exists")
        return nil
    }

    default_admin:=app.LoadDefaultAdminConfig()

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
    }

    result := db.Create(&admin)
    if result.Error != nil {
        return result.Error
    }

    log.Printf("Default admin user created with email: %s\n", adminEmail)
    return nil
}