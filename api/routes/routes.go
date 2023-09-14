package routes

import (
	"products-service/api/controllers/login"
	"products-service/api/controllers/products"
	"products-service/api/controllers/validation"
	"products-service/api/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode) //ReleaseMode, DebugMode, TestMode
	//Middlewares
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.DatabaseMiddleware(db))

	//Routes
	r.POST("/login", login.Login)
	r.GET("/validation-token", middleware.AuthMiddleware(), validation.CheckStatus)
	r.GET("/products", middleware.AuthMiddleware(), products.GetProducts)
	r.GET("/favorites", middleware.AuthMiddleware(), products.GetFavorites)
	r.POST("/favorites/:productsID", middleware.AuthMiddleware(), products.AddFavorite)
	r.DELETE("/favorites/:productsID", middleware.AuthMiddleware(), products.RemoveFavorite)

	return r
}
