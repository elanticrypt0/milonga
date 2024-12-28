package middleware

import (
    "os"
    "strings"

    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v5"
)

func Protected() fiber.Handler {
    return func(c *fiber.Ctx) error {
        authHeader := c.Get("Authorization")
        
        if authHeader == "" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "message": "unauthorized",
            })
        }

        tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
        
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("JWT_SECRET")), nil
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