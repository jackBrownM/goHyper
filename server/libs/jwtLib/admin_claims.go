package jwtLib

import "github.com/dgrijalva/jwt-go"

type AdminClaims struct {
	jwt.StandardClaims
	No string
	Id int
}

func (a AdminClaims) Valid() error {
	return a.StandardClaims.Valid()
}
