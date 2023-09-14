package validation

import "github.com/gin-gonic/gin"

func CheckStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Status OK",
	})
}
