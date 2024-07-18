package services

import (
	"dvtool/utils"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type AuthServiceInterface interface {
	ValidateCredentials(username, password string) bool
	GenerateToken(username string) (string, error)
	LoginUser(ctx *gin.Context)
}

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) ValidateCredentials(username, password string) bool {
	return username == viper.GetString("AUTH_USERNAME") && password == viper.GetString("AUTH_PASSWORD")
}

func (s *AuthService) GenerateToken(username string) (string, error) {
	return utils.GenerateToken(username)
}

func (s *AuthService) LoginUser(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if s.ValidateCredentials(username, password) {
		session := sessions.Default(ctx)
		token, err := s.GenerateToken(username)

		session.Set("token", token)
		session.Save()

		utils.CheckError(err)

		ctx.Redirect(http.StatusSeeOther, "/")
	} else {
		ctx.Redirect(http.StatusSeeOther, "/login")
	}
}
