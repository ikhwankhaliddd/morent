package services

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

var JWT_SECRET = os.Getenv("JWT_SECRET")
var SECRET_KEY = []byte(JWT_SECRET)

func (s *AuthService) GenerateToken(email string) (string, error) {
	payload := jwt.MapClaims{}
	payload["user_id"] = email

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Printf("[Auth Service] Error Generate Token with error : %v", err)
		return "", err
	}
	return signedToken, nil
}
