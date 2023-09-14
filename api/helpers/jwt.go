package helpers

import (
	models "products-service/api/models/entities"
	"products-service/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// En repository.go
func GenerateJWT(user *models.User) (string, error) {
	expirationTime := time.Now().Add(time.Duration(config.GetEnv().JWTExpirationDays) * 24 * time.Hour)
	claims := jwt.MapClaims{
		"authorized": true,
		"username":   user.Username,
		"ExpiresAt":  expirationTime.Unix(),
		"id":         user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.GetEnv().JWTSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
