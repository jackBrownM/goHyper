package jwtLib

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	jwt.StandardClaims
	Id int
}

func (a UserClaims) Valid() error {
	return a.StandardClaims.Valid()
}
