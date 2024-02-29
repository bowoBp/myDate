package middleware

import (
	"github.com/bowoBp/myDate/pkg/environment"
	"github.com/bowoBp/myDate/pkg/maker"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
)

type (
	Auth struct {
		maker  maker.Generator
		mapper mapperAuth
		env    environment.Environment
	}

	mapperAuth interface {
		GetBodyJSON(c *gin.Context) map[string]any
	}

	AuthInterface interface {
	}
)

func NewAuth() AuthInterface {
	return Auth{
		maker:  maker.DefaultMaker(),
		mapper: SharedMapper{},
		env:    environment.NewEnvironment(),
	}
}

func (receiver Auth) SignClaim(claim DefaultUserClaim) (string, error) {
	method := jwt.SigningMethodHS256
	token := &jwt.Token{
		Header: map[string]interface{}{
			"typ": "JWT",
			"alg": method.Alg(),
		},
		Claims: claim,
		Method: method,
	}
	secret := []byte(os.Getenv("SECRET"))
	tokenStr, err := token.SignedString(secret)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return tokenStr, nil
}
