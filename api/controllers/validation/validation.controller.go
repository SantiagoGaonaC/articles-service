package validation

import (
	"github.com/gin-gonic/gin"
)

func CheckStatus(c *gin.Context) {
	userID, _ := c.Get("userID")
	tokenExpiresAt, _ := c.Get("tokenExpiresAt")
	c.JSON(200, gin.H{
		"message":        "Status OK",
		"userID":         userID,
		"tokenExpiresAt": tokenExpiresAt,
	})
}
