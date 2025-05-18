package vigilante

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserAuth struct {
	ID       uuid.UUID  `gorm:"type:uuid;primary_key;"`
	Email    string     `gorm:"unique;not null"`
	Username string     `gorm:"unique;not null"`
	Password string     `gorm:"type:text;not null"`
	Role     UserRole   `gorm:"type:varchar(20);default:'user'"`
	Status   UserStatus `gorm:"type:varchar(20);default:'active'"`
	gorm.Model
}

// BeforeCreate será llamado por GORM antes de crear un nuevo usuario
func (u *UserAuth) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}

func (u *UserAuth) GetProfile(db *gorm.DB, userID string) error {
	result := db.First(u, "id = ?", userID)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *UserAuth) GetByEmail(tx *gorm.DB, email string) (*UserAuth, error) {
	userdata := &UserAuth{}
	err := tx.Model(&UserAuth{}).Where("email = ?", email).First(userdata).Error
	if err != nil {
		return nil, fmt.Errorf("error on finding user with email %s", email)
	}

	return userdata, nil
}

func (u *UserAuth) GetEnabledByEmail(tx *gorm.DB, email string) (*UserAuth, error) {
	userdata := &UserAuth{}
	err := tx.Model(&UserAuth{}).Where("email = ? AND status = ?", email, UserStatusEnabled).First(userdata).Error
	if err != nil {
		return nil, fmt.Errorf("error on finding user with email %s", email)
	}

	return userdata, nil
}

func (u *UserAuth) GetByID(tx *gorm.DB, id uuid.UUID) (*UserAuth, error) {
	userdata := &UserAuth{}
	err := tx.Model(&UserAuth{}).Where("id = ?", id).First(userdata).Error
	if err != nil {
		return nil, fmt.Errorf("error on finding user with ID %s", id)
	}

	return userdata, nil
}
