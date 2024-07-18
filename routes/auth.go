package routes

import (
	"dvtool/utils"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Controller struct{}

func Login() *Controller {
	return &Controller{}
}

func (c *Controller) Index(ctx *gin.Context) {
	sessions := sessions.Default(ctx)

	if sessions.Get("token") != nil {
		ctx.Redirect(http.StatusSeeOther, "/")
		ctx.Abort()
		return
	}

	ctx.HTML(http.StatusOK, "login", gin.H{
		"title": "Login",
	})

}

func (c *Controller) Auth(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if username == viper.GetString("AUTH_USERNAME") && password == viper.GetString("AUTH_PASSWORD") {
		session := sessions.Default(ctx)
		token, err := utils.GenerateToken(username)

		session.Set("token", token)
		session.Save()

		utils.CheckError(err)

		ctx.Redirect(http.StatusSeeOther, "/")
	} else {
		ctx.Redirect(http.StatusUnauthorized, "/login")
	}
}
