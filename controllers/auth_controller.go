package controllers

import (
	"dvtool/services"
	"dvtool/types"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService services.AuthServiceInterface
}

func NewAuthController(authService services.AuthServiceInterface) *AuthController {
	return &AuthController{AuthService: authService}
}

func (c *AuthController) Index(ctx *gin.Context) {
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

func (c *AuthController) Auth(ctx *gin.Context) {

	var authValidator types.AuthValidator

	if err := ctx.ShouldBind(&authValidator); err != nil {
		ctx.HTML(http.StatusBadRequest, "login", gin.H{
			"title": "Login",
		})

		return
	}

	c.AuthService.LoginUser(ctx)
}
