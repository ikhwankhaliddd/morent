package services

import (
	"errors"
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
	payload["user_email"] = email

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Printf("[Auth Service] Error Generate Token with error : %v", err)
		return "", err
	}
	return signedToken, nil
}

func (s *AuthService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("[Auth Service] invalid token")
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
