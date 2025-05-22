package vigilante

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// IMPORTANT: default duration
// half hour
const TOKEN_LIFETIME = 30

func CreateNewJWToken(userID uuid.UUID, email, role, jwt_secret string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * TOKEN_LIFETIME).Unix()

	t, err := token.SignedString([]byte(jwt_secret))
	if err != nil {
		return "", fmt.Errorf("could not create token")
	}

	return t, nil
}
