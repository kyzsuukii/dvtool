package routes

import (
	"dvtool/controllers"
	"dvtool/middleware"
	"dvtool/services"

	"github.com/gin-gonic/gin"
)

func WebRouter(r *gin.Engine) {
	authService := services.NewAuthService()
	authController := controllers.NewAuthController(authService)

	r.GET("/login", authController.Index)
	r.POST("/login", authController.Auth)

	r.Use(middleware.JwtVerify())

	r.GET("/", controllers.NewIndexController().Index)
}
