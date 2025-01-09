package migrations

import (
	"fmt"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	fmt.Println("Starting migration...")
	fmt.Println("Starting migration... done")
	return nil
}
