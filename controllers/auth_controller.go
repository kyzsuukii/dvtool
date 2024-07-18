package controllers

import (
	"dvtool/services"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthValidator struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type AuthController struct {
	authService services.AuthServiceInterface
}

func NewAuthController(authService services.AuthServiceInterface) *AuthController {
	return &AuthController{authService: authService}
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

	var authValidator AuthValidator

	if err := ctx.ShouldBind(&authValidator); err != nil {
		ctx.HTML(http.StatusBadRequest, "login", gin.H{
			"title": "Login",
		})

		return
	}

	c.authService.LoginUser(ctx)
}
