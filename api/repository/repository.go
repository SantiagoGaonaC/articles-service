package repository

import (
	"errors"
	"log"
	"products-service/api/models/database"
	models "products-service/api/models/entities"

	"gorm.io/gorm"
)

func ValidateUserRepo(username, password string) (*models.User, error) {
	db, _ := database.ConnectToDatabase()
	var user models.User
	if err := db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func GetAllProducts(db *gorm.DB) ([]models.ProductResponse, error) {
	var products []models.ProductResponse
	if err := db.Model(&models.Product{}).Select("id, vendor, product_name, rating, image_url").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func AddFavorite(db *gorm.DB, userID, productID uint) error {
	var user models.User
	var product models.Product
	log.Printf("userID: %d", userID)
	if err := db.First(&user, userID).Error; err != nil {
		return errors.New("user does not exist")
	}

	if err := db.First(&product, productID).Error; err != nil {
		return errors.New("product does not exist")
	}

	favorite := models.Favorite{
		UserID:    userID,
		ProductID: productID,
	}

	if err := db.Create(&favorite).Error; err != nil {
		return err
	}

	return nil
}

func RemoveFavorite(db *gorm.DB, userID, productID uint) error {
	if err := db.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&models.Favorite{}).Error; err != nil {
		return err
	}
	return nil
}

func GetFavoritesByUserID(db *gorm.DB, userID uint) ([]models.ProductResponse, error) {
	var products []models.ProductResponse
	if err := db.Joins("JOIN favorites on favorites.product_id = products.id").
		Where("favorites.user_id = ?", userID).
		Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
