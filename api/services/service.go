// service.go

package service

import (
	models "articles-service/api/models/entities"
	"articles-service/api/repository"
	"articles-service/config"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Authenticate(username, password string) (*models.LoginResponse, error) {
	user, err := repository.ValidateUserRepo(username, password)

	if err != nil {
		return nil, errors.New("Unauthorized")
	}

	// Añadir tiempo de expiración al token
	expirationTime := time.Now().Add(time.Duration(config.GetEnv().JWTExpirationDays) * 24 * time.Hour)

	claims := jwt.MapClaims{
		"authorized": true,
		"username":   username,
		"ExpiresAt":  expirationTime.Unix(),
		"ID":         fmt.Sprintf("%d", user.ID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.GetEnv().JWTSecret))
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Token:    tokenString,
	}, nil
}
