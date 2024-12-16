package jwt

import (
	"time"
	"user_service/internal/jwt"
	"user_service/internal/models"

	"github.com/golang-jwt/jwt"
)

func NewToken(user *models.User, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["phone"] = user.Phone
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenStr, err := token.SignedString()
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
