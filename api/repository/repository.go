package repository

import (
	"articles-service/api/models/database"
	models "articles-service/api/models/entities"
	"errors"
)

func ValidateUserRepo(username, password string) (*models.User, error) {
	db, _ := database.ConnectToDatabase()
	var user models.User
	if err := db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
