package vigilante

import (
	"encoding/base64"
	"fmt"
	"milonga/milonga/app"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ParseLogingByTokenInput(c *fiber.Ctx) (LoginByTokenInput, error) {
	input := LoginByTokenInput{}
	ref := c.Query("ref", "")
	if ref == "" {
		return input, fmt.Errorf("invalid credentials")
	}

	refDecoded, err := base64.StdEncoding.DecodeString(ref)
	if err != nil {
		return input, fmt.Errorf("error parsing credentials")
	}

	parts := strings.Split(string(refDecoded), ":")
	if len(parts) != 2 {
		return input, fmt.Errorf("invalid credential format")
	}

	input.Email = parts[0]
	input.PasswordToken = parts[1]
	return input, nil
}

func GenerateLoginPasswordTokenLink(app *app.App, email, passwordtoken string) string {

	refFormat := fmt.Sprintf("%s:%s", email, passwordtoken)
	ref := base64.StdEncoding.EncodeToString([]byte(refFormat))

	return fmt.Sprintf("%s/%s/auth/login/otp/link?ref=%s", app.Config.AppHost, app.Config.APIPath, ref)
}

func CreateSessionCookie(sessionName, token string) *fiber.Cookie {

	return &fiber.Cookie{
		Name:     parseSessionName(sessionName),
		Value:    token,
		Expires:  time.Now().Add(1 * time.Hour),
		Secure:   true,
		HTTPOnly: true,
		SameSite: "Lax",
	}
}

func parseSessionName(sessionName string) string {
	sessionName = strings.ToLower(sessionName)
	sessionName = strings.Replace(sessionName, "@", "AT", -1)
	return sessionName
}
