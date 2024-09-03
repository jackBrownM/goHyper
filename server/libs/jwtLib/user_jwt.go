package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/michaelzx/alc/alc_crypto"
	"net/url"
)

type UserJwt struct {
	signKey string
	aesKey  string
	claims  *UserClaims
}

func NewUserJwt(jwtSignKey, jwtAesKey string, claims *UserClaims) *UserJwt {
	return &UserJwt{
		signKey: jwtSignKey,
		aesKey:  jwtAesKey,
		claims:  claims,
	}
}
func ParseUserJwt(jwtSignKey, jwtAesKey string, aesJwtStr string) (*UserJwt, error) {
	aesKeyBytes := []byte(jwtAesKey)
	jwtBytes, err := alc_crypto.AesBase64Kit.DecryptBase64(aesJwtStr, aesKeyBytes)
	if err != nil {
		return nil, err
	}
	jwtStr := string(jwtBytes)
	token, err := jwt.ParseWithClaims(jwtStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSignKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, errors.New("claims转换失败")
	}
	return &UserJwt{
		signKey: jwtSignKey,
		aesKey:  jwtAesKey,
		claims:  claims,
	}, nil
}
func (a *UserJwt) Encode() (string, error) {
	if a.claims == nil {
		return "", errors.New("claims为初始化")
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

func (a *UserJwt) Claims() *UserClaims {
	return a.claims
}
