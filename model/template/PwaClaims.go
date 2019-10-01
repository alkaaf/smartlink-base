package template

import (
	"github.com/alkaaf/smartlink-base/model"
	"github.com/dgrijalva/jwt-go"
)

type PwaClaims struct {
	model.CustomerGlobal
	jwt.StandardClaims
}
