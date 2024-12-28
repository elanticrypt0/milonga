// api/handlers/user.go
package handlers

import (
    "errors"

    "github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
    "milonga/api/models"
	"milonga/pkg/app"
)

// GetAllUsers obtiene todos los usuarios
func GetAllUsers(c *fiber.Ctx, app *app.App) error {
    var users []models.User
    result := app.DB.Primary.Find(&users)
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
func GetUser(c *fiber.Ctx, app *app.App) error {
    id := c.Params("id")
    var user models.User

    result := app.DB.Primary.First(&user, "id = ?", id)
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
func SearchUser(c *fiber.Ctx, app *app.App) error {
    email := c.Query("email")
    username := c.Query("username")
    var user models.User

    if email == "" && username == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Email or username is required",
        })
    }

    query := app.DB.Primary
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

type UpdateUserInput struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
    Role     string `json:"role"`
}

// UpdateUser actualiza un usuario existente
func UpdateUser(c *fiber.Ctx, app *app.App) error {
    id := c.Params("id")
    input := new(UpdateUserInput)

    if err := c.BodyParser(input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Invalid input",
        })
    }

    var user models.User
    result := app.DB.Primary.First(&user, "id = ?", id)
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
    if input.Role != "" {
        // Solo los admins pueden cambiar roles
        if tokenUser["role"] != "admin" {
            return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
                "message": "Not authorized to change roles",
            })
        }
        updates["role"] = input.Role
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

    result = app.DB.Primary.Model(&user).Updates(updates)
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
func DeleteUser(c *fiber.Ctx, app *app.App) error {
    id := c.Params("id")
    
    // Verificar si el usuario tiene permisos para eliminar
    tokenUser := c.Locals("user").(jwt.MapClaims)
    if tokenUser["role"] != "admin" {
        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
            "message": "Not authorized to delete users",
        })
    }

    var user models.User
    result := app.DB.Primary.First(&user, "id = ?", id)
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

    result = app.DB.Primary.Delete(&user)
    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Error deleting user",
        })
    }

    return c.JSON(fiber.Map{
        "message": "User deleted successfully",
    })
}