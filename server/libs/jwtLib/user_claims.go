package jwtLib

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	jwt.StandardClaims
	No string
}

func (a UserClaims) Valid() error {
	return a.StandardClaims.Valid()
}
