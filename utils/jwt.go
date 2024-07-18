package utils

import (
	"github.com/cristalhq/jwt/v5"
	"github.com/gofrs/uuid"
	"github.com/spf13/viper"
)

type UserClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

func GenerateToken(username string) (string, error) {
	key := []byte(viper.GetString("JWT_SECRET"))
	uuidv4, err := uuid.NewV4()

	CheckError(err)

	signer, err := jwt.NewSignerHS(jwt.HS256, key)

	CheckError(err)

	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Audience: []string{"admin"},
			ID:       uuidv4.String(),
		},
		Username: username,
	}

	buider := jwt.NewBuilder(signer)

	token, err := buider.Build(claims)

	return token.String(), err
}
