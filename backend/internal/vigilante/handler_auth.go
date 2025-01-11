package vigilante

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"milonga/internal/app"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	app *app.App
	db  *gorm.DB
}

func NewAuthHandler(app *app.App, db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		app: app,
		db:  db,
	}
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginByTokenInput struct {
	Email         string `json:"email"`
	PasswordToken string `json:"token"`
}

type RegisterInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (me *AuthHandler) Register(c *fiber.Ctx) error {
	input := new(RegisterInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error hashing password",
		})
	}

	user := &User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	result := me.db.Create(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

func (me *AuthHandler) Login(c *fiber.Ctx) error {
	input := new(LoginInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}

	var user User
	result := me.db.Where("email = ? AND status = ?", input.Email, UserStatusEnabled).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	// Verify password
	err := ComparePassword(user.Password, input.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	t, err := CreateNewToken(user.ID, user.Email, string(user.Role), me.app.Config.JWTSecret)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	return c.JSON(fiber.Map{
		"token": t,
		"user": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

func (me *AuthHandler) LoginByPasswordToken(c *fiber.Ctx) error {
	input := new(LoginByTokenInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}

	var user User
	result := me.db.Where("email = ? AND status = ?", input.Email, UserStatusEnabled).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	passToken := NewPasswordToken()

	err := passToken.CheckToken(user.ID, input.PasswordToken, me.db)
	if err != nil {

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			// "message": "Invalid credentials - token",
			"message": fmt.Sprintf("%s", err),
		})
	}

	t, err := CreateNewToken(user.ID, user.Email, string(user.Role), me.app.Config.JWTSecret)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	return c.JSON(fiber.Map{
		"token": t,
		"user": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

// GetProfile obtiene la información del usuario autenticado
func (me *AuthHandler) GetProfile(c *fiber.Ctx) error {
	// Obtener los claims del token JWT
	tokenUser := c.Locals("user").(jwt.MapClaims)
	userID := tokenUser["user_id"].(string)

	user := &User{}
	err := user.GetProfile(me.db, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting user profile",
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
