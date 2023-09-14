package routes

import (
	"articles-service/api/controllers/articles"
	"articles-service/api/controllers/login"
	"articles-service/api/controllers/validation"
	"articles-service/api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode) //ReleaseMode, DebugMode, TestMode
	//Middlewares
	r.Use(middleware.CORSMiddleware())

	//Routes
	r.POST("/login", login.Login)
	r.GET("/validation-token", middleware.AuthMiddleware(), validation.CheckStatus)
	r.GET("/articles", middleware.AuthMiddleware(), articles.GetArticles)

	return r
}
