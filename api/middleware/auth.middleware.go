package middleware

import (
	"errors"
	"net/http"
	"products-service/config"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
			c.Set("userID", uint(userID))
			c.Set("username", claims["username"])
		} else {
			c.JSON(http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		c.Next()
	}
}
