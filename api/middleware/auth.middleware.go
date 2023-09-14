package middleware

import (
	"errors"
	"net/http"
	models "products-service/api/models/entities"
	"products-service/config"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) != 2 {
			c.JSON(http.StatusUnauthorized, "Malformed token")
			c.Abort()
			return
		}

		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, errors.New("invalid signing method")
			}
			return []byte(config.GetEnv().JWTSecret), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID, ok := claims["id"].(float64)
			if !ok {
				c.JSON(http.StatusUnauthorized, "Invalid token")
				c.Abort()
				return
			}

			db := c.MustGet("db").(*gorm.DB)
			var user models.User
			if err := db.First(&user, uint(userID)).Error; err != nil {
				c.JSON(http.StatusUnauthorized, "User not found")
				c.Abort()
				return
			}

			// expirado?
			exp, ok := claims["ExpiresAt"].(float64)
			if !ok || time.Unix(int64(exp), 0).Before(time.Now()) {
				c.JSON(http.StatusUnauthorized, "Token has expired")
				c.Abort()
				return
			}

			c.Set("userID", uint(userID))
			c.Set("username", claims["username"])
			c.Set("tokenExpiresAt", exp)
		} else {
			c.JSON(http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		c.Next()
	}
}
