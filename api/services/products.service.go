// service.go

package service

import (
	models "products-service/api/models/entities"
	"products-service/api/repository"

	"gorm.io/gorm"
)

func GetAllProducts(db *gorm.DB) ([]models.Product, error) {
	return repository.GetAllProducts(db)
}

func AddFavorite(db *gorm.DB, userID, productID uint) error {
	return repository.AddFavorite(db, userID, productID)
}

func RemoveFavorite(db *gorm.DB, userID, productID uint) error {
	return repository.RemoveFavorite(db, userID, productID)
}

func GetFavoritesByUserID(db *gorm.DB, userID uint) ([]models.Product, error) {
	return repository.GetFavoritesByUserID(db, userID)
}
