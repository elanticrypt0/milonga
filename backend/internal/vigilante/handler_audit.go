package vigilante

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (me *AuthHandler) Login_audit(c *fiber.Ctx) error {
	input := new(LoginInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}

	var user User
	result := me.db.Where("email = ? AND status = ?", input.Email, UserStatusEnabled).First(&user)
	if result.Error != nil {
		// Registrar intento fallido
		loginAudit := NewLoginAudit()
		loginAudit.RegisterFailedLogin(
			uuid.Nil, // UUID nulo porque el usuario no existe
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

	// Verificar password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		// Registrar intento fallido
		loginAudit := NewLoginAudit()
		loginAudit.RegisterFailedLogin(
			user.ID,
			c.IP(),
			c.Get("User-Agent"),
			LoginMethodPassword,
			"Invalid password",
			me.db,
		)

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	// Generar JWT
	t, err := CreateNewToken(user.ID, user.Email, string(user.Role), me.app.Config.JWTSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	// Registrar login exitoso
	loginAudit := NewLoginAudit()
	loginAudit.RegisterSuccessfulLogin(
		user.ID,
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

func (me *AuthHandler) LoginByPasswordToken_audit(c *fiber.Ctx) error {
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

		// Registrar intento fallido
		loginAudit := NewLoginAudit()
		loginAudit.RegisterFailedLogin(
			user.ID,
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

	t, err := CreateNewToken(user.ID, user.Email, string(user.Role), me.app.Config.JWTSecret)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	// Registrar login exitoso
	loginAudit := NewLoginAudit()
	loginAudit.RegisterSuccessfulLogin(
		user.ID,
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
