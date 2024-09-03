package jwt

import "github.com/dgrijalva/jwt-go"

type AdminClaims struct {
	jwt.StandardClaims
	No string
}

func (a AdminClaims) Valid() error {
	return a.StandardClaims.Valid()
}
