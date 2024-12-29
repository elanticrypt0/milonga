package vigilante

import (
	"errors"
	"fmt"

	"milonga/internal/app"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	app *app.App
	db  *gorm.DB
}

func NewUserHandler(app *app.App, db *gorm.DB) *UserHandler {
	return &UserHandler{
		app: app,
		db:  db,
	}
}

// GetAllUsers obtiene todos los usuarios
func (me *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	var users []User
	result := me.db.Find(&users)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting users",
		})
	}

	// Limpiamos las contraseñas antes de enviar
	var cleanUsers []fiber.Map
	for _, user := range users {
		cleanUsers = append(cleanUsers, fiber.Map{
			"id":        user.ID,
			"username":  user.Username,
			"email":     user.Email,
			"role":      user.Role,
			"createdAt": user.CreatedAt,
			"updatedAt": user.UpdatedAt,
		})
	}

	return c.JSON(fiber.Map{
		"users": cleanUsers,
	})
}

// GetUser obtiene un usuario por ID
func (me *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User

	result := me.db.First(&user, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting user",
		})
	}

	return c.JSON(fiber.Map{
		"user": fiber.Map{
			"id":        user.ID,
			"username":  user.Username,
			"email":     user.Email,
			"role":      user.Role,
			"createdAt": user.CreatedAt,
			"updatedAt": user.UpdatedAt,
		},
	})
}

// SearchUser busca un usuario por email o username
func (me *UserHandler) SearchUser(c *fiber.Ctx) error {
	email := c.Query("email")
	username := c.Query("username")
	var user User

	if email == "" && username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email or username is required",
		})
	}

	query := me.db
	if email != "" {
		query = query.Where("email = ?", email)
	}
	if username != "" {
		query = query.Or("username = ?", username)
	}

	result := query.First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error searching user",
		})
	}

	return c.JSON(fiber.Map{
		"user": fiber.Map{
			"id":        user.ID,
			"username":  user.Username,
			"email":     user.Email,
			"role":      user.Role,
			"createdAt": user.CreatedAt,
			"updatedAt": user.UpdatedAt,
		},
	})
}

// CreateUserInput define la estructura para crear un nuevo usuario
type CreateUserInput struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role" validate:"required,oneof=user admin"`
}

// CreateUser crea un nuevo usuario (solo admins)
func (me *UserHandler) CreateUser(c *fiber.Ctx) error {
	// Verificar que el usuario sea admin
	tokenUser := c.Locals("user").(jwt.MapClaims)

	if tokenUser["role"] != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Not authorized to create users",
		})
	}

	input := new(CreateUserInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}

	// Verificar si ya existe un usuario con ese email o username
	var existingUser User
	if result := me.db.Where("email = ? OR username = ?", input.Email, input.Username).First(&existingUser); result.Error == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "User with this email or username already exists",
		})
	}

	// Hash de la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error hashing password",
		})
	}

	user_role, err := NewUserRole(input.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	// Crear el nuevo usuario
	user := &User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     user_role,
	}

	result := me.db.Create(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating user",
			"error":   result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user": fiber.Map{
			"id":        user.ID,
			"username":  user.Username,
			"email":     user.Email,
			"role":      user.Role,
			"createdAt": user.CreatedAt,
			"updatedAt": user.UpdatedAt,
		},
	})
}

type UpdateUserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// UpdateUser actualiza un usuario existente
func (me *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	input := new(UpdateUserInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}

	var user User
	result := me.db.First(&user, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting user",
		})
	}

	// Verificar si el usuario tiene permisos para actualizar
	tokenUser := c.Locals("user").(jwt.MapClaims)
	if tokenUser["user_id"] != id && tokenUser["role"] != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Not authorized to update this user",
		})
	}

	updates := make(map[string]interface{})
	if input.Username != "" {
		updates["username"] = input.Username
	}
	if input.Email != "" {
		updates["email"] = input.Email
	}

	user_role, err := NewUserRole(input.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	if input.Role != "" {
		// Solo los admins pueden cambiar roles
		if tokenUser["role"] != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Not authorized to change roles",
			})
		}
		updates["role"] = user_role
	}
	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error hashing password",
			})
		}
		updates["password"] = string(hashedPassword)
	}

	result = me.db.Model(&user).Updates(updates)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error updating user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User updated successfully",
		"user": fiber.Map{
			"id":        user.ID,
			"username":  user.Username,
			"email":     user.Email,
			"role":      user.Role,
			"createdAt": user.CreatedAt,
			"updatedAt": user.UpdatedAt,
		},
	})
}

// DeleteUser elimina un usuario
func (me *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	// Verificar si el usuario tiene permisos para eliminar
	tokenUser := c.Locals("user").(jwt.MapClaims)

	token_role := fmt.Sprintf("%v", tokenUser["role"])

	if IsAdmin(token_role) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Not authorized to delete users",
		})
	}

	var user User
	result := me.db.First(&user, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting user",
		})
	}

	result = me.db.Delete(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error deleting user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
