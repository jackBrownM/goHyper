package jwtLib

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/michaelzx/alc/alc_crypto"
	"net/url"
)

type AdminJwt struct {
	signKey string
	aesKey  string
	claims  *AdminClaims
}

func NewAdminJwt(jwtSignKey, jwtAesKey string, claims *AdminClaims) *AdminJwt {
	return &AdminJwt{
		signKey: jwtSignKey,
		aesKey:  jwtAesKey,
		claims:  claims,
	}
}
func DecodeAdminJwt(jwtSignKey, jwtAesKey string, jwtString string) (*AdminJwt, error) {
	aesJwtStr, err := url.QueryUnescape(jwtString)
	if err != nil {
		return nil, err
	}
	aesKeyBytes := []byte(jwtAesKey)
	jwtBytes, err := alc_crypto.AesBase64Kit.DecryptBase64(aesJwtStr, aesKeyBytes)
	if err != nil {
		return nil, err
	}
	jwtStr := string(jwtBytes)
	token, err := jwt.ParseWithClaims(jwtStr, &AdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSignKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*AdminClaims)
	if !ok {
		return nil, errors.New("claims转换失败")
	}
	return &AdminJwt{
		signKey: jwtSignKey,
		aesKey:  jwtAesKey,
		claims:  claims,
	}, nil
}
func (a *AdminJwt) Encode() (string, error) {
	if a.claims == nil {
		return "", errors.New("claims未初始化")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a.claims)
	if str, err := token.SignedString([]byte(a.signKey)); err == nil {
		jwtBytes := []byte(str)
		aesKeyBytes := []byte(a.aesKey)
		aesJwt := alc_crypto.AesBase64Kit.EncryptBase64(jwtBytes, aesKeyBytes)
		aesJwt = url.QueryEscape(aesJwt)
		return aesJwt, nil
	} else {
		return "", err
	}
}

func (a *AdminJwt) Claims() *AdminClaims {
	return a.claims
}
