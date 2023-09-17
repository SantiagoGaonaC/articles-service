package validation

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckStatus(c *gin.Context) {
	userID, _ := c.Get("userID")
	tokenExpiresAt, _ := c.Get("tokenExpiresAt")

	authorizationHeader := c.GetHeader("Authorization")
	parts := strings.SplitN(authorizationHeader, " ", 2)
	var token string
	if len(parts) == 2 {
		token = parts[1]
	}

	c.JSON(200, gin.H{
		"id":             userID,
		"username":       "admin",
		"statusCode":     http.StatusOK,
		"message":        "Status OK",
		"token":          token,
		"tokenExpiresAt": tokenExpiresAt,
	})
}
