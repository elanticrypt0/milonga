package models

import (
	"database/sql/driver"
	"fmt"
)

type UserRole string

const (
	UserRoleAdmin      UserRole = "admin"      // Acceso total al sistema
	UserRoleSupervisor UserRole = "supervisor" // Puede supervisar moderadores y usuarios
	UserRoleModerator  UserRole = "moderator"  // Modera contenido y usuarios
	UserRoleAnalyst    UserRole = "analyst"    // Acceso a análisis y reportes
	UserRoleSupport    UserRole = "support"    // Soporte técnico y atención al usuario
	UserRoleUser       UserRole = "user"       // Usuario regular
)

var roleMap = map[string]UserRole{
	"admin":      UserRoleAdmin,
	"supervisor": UserRoleSupervisor,
	"moderator":  UserRoleModerator,
	"analyst":    UserRoleAnalyst,
	"support":    UserRoleSupport,
	"user":       UserRoleUser,
}

// comprueba si un rol de usuario es más alto que el segundo.
func IsHigherRole(role1, role2 string) bool {
	roleWeight := map[UserRole]int{
		UserRoleAdmin:      60,
		UserRoleSupervisor: 50,
		UserRoleModerator:  40,
		UserRoleAnalyst:    30,
		UserRoleSupport:    20,
		UserRoleUser:       10,
	}
	return roleWeight[roleMap[role1]] > roleWeight[roleMap[role2]]
}

func NewUserRole(str string) (UserRole, error) {
	if role, exists := roleMap[str]; exists {
		return role, nil
	}
	return "", fmt.Errorf("rol desconocido: %s", str)
}

func (r UserRole) IsValid() bool {
	switch r {
	case UserRoleAdmin, UserRoleSupervisor, UserRoleModerator,
		UserRoleAnalyst, UserRoleSupport, UserRoleUser:
		return true
	}
	return false
}

func (r UserRole) String() string {
	return string(r)
}

func (r *UserRole) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		*r = UserRole(string(v))
	case string:
		*r = UserRole(v)
	default:
		return fmt.Errorf("unsupported type for UserRole: %T", value)
	}

	if !r.IsValid() {
		return fmt.Errorf("invalid role value: %s", *r)
	}

	return nil
}

func (r UserRole) Value() (driver.Value, error) {
	if !r.IsValid() {
		return nil, fmt.Errorf("invalid role value: %s", r)
	}
	return string(r), nil
}

// Métodos auxiliares para verificar permisos
func (r UserRole) CanModerate() bool {
	switch r {
	case UserRoleAdmin, UserRoleSupervisor, UserRoleModerator:
		return true
	default:
		return false
	}
}

func (r UserRole) CanAccessAnalytics() bool {
	switch r {
	case UserRoleAdmin, UserRoleSupervisor, UserRoleAnalyst:
		return true
	default:
		return false
	}
}

func (r UserRole) CanManageUsers() bool {
	return r == UserRoleAdmin || r == UserRoleSupervisor
}

func (r UserRole) CanAccessSupportTools() bool {
	switch r {
	case UserRoleAdmin, UserRoleSupervisor, UserRoleSupport:
		return true
	default:
		return false
	}
}

func IsAdmin(user_role string) bool {

	if !IsHigherRole("admin", user_role) {
		return true
	} else {
		return false
	}
}

func IsUser(user_role string) bool {

	if !IsHigherRole(user_role, "user") {
		return true
	} else {
		return false
	}
}
