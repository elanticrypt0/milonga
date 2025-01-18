package vigilante

import (
	"fmt"
	"milonga/internal/app"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type VigilanteMiddleware struct {
	app       *app.App
	jwtSecret string
}

func NewVigilanteMiddelware(app *app.App) *VigilanteMiddleware {
	return &VigilanteMiddleware{
		app:       app,
		jwtSecret: app.Config.JWTSecret,
	}
}

// ValidateToken es una función auxiliar que valida el token y retorna los claims
func (me *VigilanteMiddleware) ValidateToken(c *fiber.Ctx) (jwt.MapClaims, error) {

	// auth be header
	authHeader := c.Get("Authorization")

	if authHeader != "" {
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := me.isTokenValid(tokenString)
		if err != nil {
			return nil, err
		}
		return token, nil
	}

	// auth by cookies
	tokenCookie := c.Cookies("userSession", "x")
	if tokenCookie != "x" {
		token, err := me.isTokenValid(tokenCookie)
		if err != nil {
			return nil, err
		}
		return token, nil
	}

	return nil, fmt.Errorf("no authorization")

}

// Protected verifica que el token sea válido
func (me *VigilanteMiddleware) IsLogged() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, err := me.ValidateToken(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "unauthorized",
			})
		}

		c.Locals("user", claims)
		return c.Next()
	}
}

// RequireRole verifica que el usuario tenga un rol específico
func (me *VigilanteMiddleware) RequireRole(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals("user").(jwt.MapClaims)
		userRole := fmt.Sprintf("%v", claims["role"])

		// Verificar si el rol del usuario está en los roles permitidos
		for _, role := range allowedRoles {
			if userRole == role {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}
}

// IsStaff verifica que el usuario no tenga rol de usuario
func (me *VigilanteMiddleware) IsStaff() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals("user").(jwt.MapClaims)

		if IsUser(fmt.Sprintf("%v", claims["role"])) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "unauthorized",
			})
		}
		return c.Next()
	}
}

func (me *VigilanteMiddleware) isTokenValid(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(me.jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}

// IsSameUserAsQuery verifica que el usuario que solicite información de un usuario sea el mismo
// esta verificación se salta si es parte del staff
func (me *VigilanteMiddleware) IsSameUserAsQuery() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		claims := c.Locals("user").(jwt.MapClaims)

		if IsUser(fmt.Sprintf("%v", claims["role"])) {
			if id != fmt.Sprintf("%v", claims["id"]) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "unauthorized",
				})

			}
			return c.Next()

		}

		return c.Next()

	}
}
