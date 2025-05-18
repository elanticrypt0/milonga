package vigilante

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"milonga/milonga/app"

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

	hashedPasswordStr := string(hashedPassword)

	user := &UserAuth{
		Username: input.Username,
		Email:    input.Email,
		Password: hashedPasswordStr,
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

	var user UserAuth
	result := me.db.Where("email = ? AND status = ?", input.Email, UserStatusEnabled).First(&user)
	if result.Error != nil {

		// Registrar intento fallido
		loginAudit := NewLoginAudit()
		loginAudit.RegisterFailedLogin(
			user.ID, // UUID nulo porque el usuario no existe
			input.Email,
			c.IP(),
			c.Get("User-Agent"),
			LoginMethodPassword,
			"User not found",
			me.db,
		)

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	// Verify password
	err := ComparePassword(user.Password, input.Password)
	if err != nil {

		// Registrar intento fallido
		loginAudit := NewLoginAudit()
		loginAudit.RegisterFailedLogin(
			user.ID, // UUID nulo porque el usuario no existe
			input.Email,
			c.IP(),
			c.Get("User-Agent"),
			LoginMethodPassword,
			"Invalid password",
			me.db,
		)

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials -2",
		})
	}

	t, err := CreateNewJWToken(user.ID, user.Email, string(user.Role), me.app.Config.JWTSecret)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	// Registrar login exitoso
	loginAudit := NewLoginAudit()
	loginAudit.RegisterSuccessfulLogin(
		user.ID,
		input.Email,
		c.IP(),
		c.Get("User-Agent"),
		LoginMethodPassword,
		me.db,
	)

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

	user := &UserAuth{}
	user, err := user.GetEnabledByEmail(me.db, input.Email)
	if err != nil {
		// Registrar intento fallido
		loginAudit := NewLoginAudit()
		loginAudit.RegisterFailedLogin(
			user.ID,
			input.Email,
			c.IP(),
			c.Get("User-Agent"),
			LoginMethodToken,
			"User not found",
			me.db,
		)

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	passToken := NewPasswordToken()

	err = passToken.CheckToken(user.ID, input.PasswordToken, me.db)
	if err != nil {

		// Registrar intento fallido
		loginAudit := NewLoginAudit()
		loginAudit.RegisterFailedLogin(
			user.ID,
			input.Email,
			c.IP(),
			c.Get("User-Agent"),
			LoginMethodToken,
			"Invalid token",
			me.db,
		)

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	// jwt
	t, err := CreateNewJWToken(user.ID, user.Email, string(user.Role), me.app.Config.JWTSecret)

	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	// Registrar login exitoso
	loginAudit := NewLoginAudit()
	loginAudit.RegisterSuccessfulLogin(
		user.ID,
		input.Email,
		c.IP(),
		c.Get("User-Agent"),
		LoginMethodToken,
		me.db,
	)

	return c.JSON(fiber.Map{
		"token": t,
		"user": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

func (me *AuthHandler) LoginByPasswordTokenWithLink(c *fiber.Ctx) error {
	input, err := ParseLogingByTokenInput(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": fmt.Sprintf("%s", err),
		})
	}

	var user UserAuth
	result := me.db.Where("email = ? AND status = ?", input.Email, UserStatusEnabled).First(&user)
	if result.Error != nil {

		// Registrar intento fallido
		loginAudit := NewLoginAudit()
		loginAudit.RegisterFailedLogin(
			user.ID,
			input.Email,
			c.IP(),
			c.Get("User-Agent"),
			LoginMethodToken,
			"User not found",
			me.db,
		)

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	passToken := NewPasswordToken()

	err = passToken.CheckToken(user.ID, input.PasswordToken, me.db)
	if err != nil {

		// Registrar intento fallido
		loginAudit := NewLoginAudit()
		loginAudit.RegisterFailedLogin(
			user.ID,
			input.Email,
			c.IP(),
			c.Get("User-Agent"),
			LoginMethodToken,
			"Invalid token",
			me.db,
		)

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	t, err := CreateNewJWToken(user.ID, user.Email, string(user.Role), me.app.Config.JWTSecret)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	// Registrar login exitoso
	loginAudit := NewLoginAudit()
	loginAudit.RegisterSuccessfulLogin(
		user.ID,
		input.Email,
		c.IP(),
		c.Get("User-Agent"),
		LoginMethodToken,
		me.db,
	)

	return c.JSON(fiber.Map{
		"token": t,
		"user": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

func (me *AuthHandler) Logout(c *fiber.Ctx) error {

	tokenUser := c.Locals("user").(jwt.MapClaims)
	userID := tokenUser["user_id"].(string)

	// TODO

	// register logout
	loginAudit := NewLoginAudit()
	loginAudit.RegisterLogout(
		userID,
		me.db,
	)

	return nil

}

// GetProfile obtiene la información del usuario autenticado
func (me *AuthHandler) GetProfile(c *fiber.Ctx) error {
	// Obtener los claims del token JWT
	tokenUser := c.Locals("user").(jwt.MapClaims)
	userID := tokenUser["user_id"].(string)

	user := &UserAuth{}
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
