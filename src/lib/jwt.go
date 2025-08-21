package lib

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/psuarezdev/go-api-starter/src/user"
)

func GenerateToken(user *user.User) (string, error) {
	godotenv.Load()
	secretKey := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"iat":      jwt.NewNumericDate(time.Now()),
		"exp":      jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ValidateToken(tokenString string) int {
	godotenv.Load()
	secretKey := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return -1
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return -1
	}

	userIdFloat, ok := claims["id"].(float64)
	if !ok {
		return -1
	}

	return int(userIdFloat)
}
