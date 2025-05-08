package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type AuthService struct {
	SecretKey string
}

func (s *AuthService) GenerateJWT(email, password string) (string, error) {

	if email != "test@example.com" || password != "1234" {
		return "", errors.New("invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 5).Unix(),
	})

	return token.SignedString([]byte(s.SecretKey))
}
