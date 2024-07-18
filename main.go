package main

import (
	"dvtool/config"
	"dvtool/middleware"
	"dvtool/routes"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	config.LoadEnv()

	r := gin.Default()

	store := cookie.NewStore([]byte(viper.GetString("SESSION_SECRET")))

	r.Use(sessions.Sessions("dvtool", store))

	r.HTMLRender = config.ViewRenderer()

	r.SetTrustedProxies(nil)

	r.Static("/assets", "./assets")

	r.GET("/login", routes.Login().Index)
	r.POST("/login", routes.Login().Auth)

	r.Use(middleware.JwtVerify())

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index", gin.H{
			"title": "Home",
		})
	})

	r.Run(":3000")
}
