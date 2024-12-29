package models

import (
	"database/sql/driver"
	"fmt"
)

// UserStatus es un tipo personalizado para manejar los estados del usuario
type UserStatus string

// Constantes para los posibles estados del usuario
const (
	UserStatusEnabled  UserStatus = "active"
	UserStatusDisabled UserStatus = "innactive"
	UserStatusPending  UserStatus = "pending"
	UserStatusBlocked  UserStatus = "blocked"
)

// Validación de estados válidos
func (s UserStatus) IsValid() bool {
	switch s {
	case UserStatusEnabled, UserStatusDisabled, UserStatusPending, UserStatusBlocked:
		return true
	}
	return false
}

// String implementa la interface Stringer
func (s UserStatus) String() string {
	return string(s)
}

// Scan implementa la interface sql.Scanner
func (s *UserStatus) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		*s = UserStatus(string(v))
	case string:
		*s = UserStatus(v)
	default:
		return fmt.Errorf("unsupported type for UserStatus: %T", value)
	}

	if !s.IsValid() {
		return fmt.Errorf("invalid status value: %s", *s)
	}

	return nil
}

// Value implementa la interface driver.Valuer
func (s UserStatus) Value() (driver.Value, error) {
	if !s.IsValid() {
		return nil, fmt.Errorf("invalid status value: %s", s)
	}
	return string(s), nil
}
