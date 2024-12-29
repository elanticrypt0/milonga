package middleware

import (
	"fmt"
	"milonga/api/models"
	"milonga/pkg/app"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// ValidateToken es una función auxiliar que valida el token y retorna los claims
func ValidateToken(c *fiber.Ctx, jwtSecret string) (jwt.MapClaims, error) {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return nil, fmt.Errorf("no authorization header")
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}

// Protected verifica que el token sea válido
func IsLogged(app *app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, err := ValidateToken(c, app.Config.JWTSecret)
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
func RequireRole(app *app.App, allowedRoles ...string) fiber.Handler {
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

// NotUser verifica que el usuario no tenga rol de usuario
func NotUser(app *app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals("user").(jwt.MapClaims)
		if models.IsUser(fmt.Sprintf("%v", claims["role"])) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "unauthorized",
			})
		}
		return c.Next()
	}
}
