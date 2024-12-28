package middleware

import (
    "strings"

    "milonga/pkg/app"
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v5"
)

func Protected(app *app.App) fiber.Handler {
    return func(c *fiber.Ctx) error {
        authHeader := c.Get("Authorization")
        
        if authHeader == "" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "message": "unauthorized",
            })
        }

        tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
        
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(app.Config.JWTSecret), nil
        })

        if err != nil || !token.Valid {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "message": "invalid token",
            })
        }

        claims := token.Claims.(jwt.MapClaims)
        c.Locals("user", claims)
        
        return c.Next()
    }
}