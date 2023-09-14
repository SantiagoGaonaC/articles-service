package service

import (
	"errors"
	"products-service/api/helpers"
	models "products-service/api/models/entities"
	"products-service/api/repository"
)

func Authenticate(username, password string) (*models.LoginResponse, error) {
	user, err := repository.ValidateUserRepo(username, password)
	if err != nil {
		return nil, errors.New("Unauthorized")
	}

	tokenString, err := helpers.GenerateJWT(user)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Token:    tokenString,
	}, nil
}
