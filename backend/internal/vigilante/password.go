package vigilante

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plainpassword string) (string, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainpassword), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password")
	}

	return string(hashedPassword), nil
}

func ComparePassword(storedPassword, inputPassword string) error {

	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(inputPassword))
	if err != nil {
		return fmt.Errorf("passwords not match")
	}
	return nil
}
