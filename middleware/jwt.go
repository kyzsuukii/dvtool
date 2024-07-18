package middleware

import (
	"dvtool/utils"

	"github.com/cristalhq/jwt/v5"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func JwtVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		secret := []byte(viper.GetString("JWT_SECRET"))

		token := session.Get("token")

		if token == nil {
			ctx.Redirect(302, "/login")
			ctx.Abort()
			return
		}

		verifier, err := jwt.NewVerifierHS(jwt.HS256, secret)

		utils.CheckError(err)

		newToken, err := jwt.Parse([]byte(token.(string)), verifier)

		utils.CheckError(err)

		err = verifier.Verify(newToken)

		if err != nil {
			ctx.Redirect(302, "/login")
			ctx.Abort()

			return
		}

		ctx.Next()
	}
}
