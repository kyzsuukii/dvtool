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

	actionService := services.NewActionService()
	actionController := controllers.NewActionController(actionService)

	r.GET("/", actionController.Index)
	r.POST("/", actionController.Output)
}
