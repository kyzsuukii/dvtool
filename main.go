package main

import (
	"dvtool/config"
	"dvtool/routes"

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

	routes.WebRouter(r)

	r.Run(":3000")
}
