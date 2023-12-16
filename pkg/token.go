package pkg

import (
	"os"
	"time"

	"github.com/DanilaNik/IU5_RIP2023/internal/service/role"
	"github.com/golang-jwt/jwt"
)

func GenerateJWTToken(userID uint, role role.Role) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID
	claims["role"] = role
	claims["exp"] = time.Hour * 1

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
