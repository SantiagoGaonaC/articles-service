package login

import (
	"net/http"
	models "products-service/api/models/entities"
	service "products-service/api/services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var request models.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	result, err := service.Authenticate(request.Username, request.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
