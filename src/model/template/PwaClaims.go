package template

import (
	"github.com/dgrijalva/jwt-go"
	"model"
)

type PwaClaims struct {
	model.CustomerGlobal
	jwt.StandardClaims
}
