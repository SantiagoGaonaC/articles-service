package products

import (
	"net/http"
	service "products-service/api/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetFavorites(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve user ID"})
		return
	}
	actualUserID := userID.(uint)
	favorites, err := service.GetFavoritesByUserID(db, uint(actualUserID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, favorites)
}

func AddFavorite(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve user ID"})
		return
	}
	actualUserID := userID.(uint)
	productsID, err := strconv.Atoi(c.Param("productsID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid products ID"})
		return
	}
	if err := service.AddFavorite(db, uint(actualUserID), uint(productsID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	favorites, err := service.GetFavoritesByUserID(db, uint(actualUserID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, favorites)
}

func RemoveFavorite(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve user ID"})
		return
	}
	actualUserID := userID.(uint)
	productsID, err := strconv.Atoi(c.Param("productsID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid products ID"})
		return
	}
	if err := service.RemoveFavorite(db, uint(actualUserID), uint(productsID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	favorites, err := service.GetFavoritesByUserID(db, uint(actualUserID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, favorites)
}
