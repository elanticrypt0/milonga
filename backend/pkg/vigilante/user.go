package vigilante

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
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
