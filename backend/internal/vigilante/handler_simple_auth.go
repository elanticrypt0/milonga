package vigilante

import (
	"fmt"

	"gorm.io/gorm"

	"milonga/internal/app"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type SimpleAuthHandler struct {
	app *app.App
	db  *gorm.DB
}

func NewSimpleAuthHandler(app *app.App, db *gorm.DB) *SimpleAuthHandler {
	return &SimpleAuthHandler{
		app: app,
		db:  db,
	}
}

type SimpleLoginInput struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (me *SimpleAuthHandler) Login(c *fiber.Ctx) error {
	input := new(SimpleLoginInput)

	if err := c.BodyParser(input); err != nil {
		return fmt.Errorf("invalid input")
	}

	var user User
	result := me.db.Where("email = ? AND status = ?", input.Email, UserStatusEnabled).First(&user)
	if result.Error != nil {

		// Registrar intento fallido
		loginAudit := NewLoginAudit()
		loginAudit.RegisterFailedLogin(
			uuid.Nil, // UUID nulo porque el usuario no existe
			input.Email,
			c.IP(),
			c.Get("User-Agent"),
			LoginMethodSimple,
			"User not found",
			me.db,
		)

		return fmt.Errorf("invalid credentials")
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
			LoginMethodSimple,
			"Invalid password",
			me.db,
		)

		return fmt.Errorf("invalid password")
	}

	t, err := CreateNewJWToken(user.ID, user.Email, string(user.Role), me.app.Config.JWTSecret)

	if err != nil {
		return fmt.Errorf("could not login")
	}

	// Registrar login exitoso
	loginAudit := NewLoginAudit()
	loginAudit.RegisterSuccessfulLogin(
		user.ID,
		input.Email,
		c.IP(),
		c.Get("User-Agent"),
		LoginMethodSimple,
		me.db,
	)

	c.Cookie(CreateSessionCookie(t))

	return nil
}
