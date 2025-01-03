package jwt

import (
	"fmt"
	"time"
	"user_service/internal/models"

	"github.com/golang-jwt/jwt"
)

func NewToken(user *models.User, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["phone"] = user.Phone
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenStr, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", fmt.Errorf("get token: %w", err)
	}

	return tokenStr, nil
}
