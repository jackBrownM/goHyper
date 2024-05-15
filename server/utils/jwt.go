package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CustomClaims 结构体用于存储自定义数据
type CustomClaims struct {
	UserID int `json:"userID"`
	jwt.StandardClaims
}

// JWTService 包含JWT操作的方法
type JWTService struct {
	SecretKey []byte // 秘钥
}

// NewJWTService 初始化JWT服务
func NewJWTService(secret string) *JWTService {
	return &JWTService{
		SecretKey: []byte(secret),
	}
}

// GenerateToken 创建JWT
func (j *JWTService) GenerateToken(userID int) (string, error) {
	claims := &CustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 过期时间为24小时
			Issuer:    "your-app-name",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(j.SecretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateToken 验证JWT
func (j *JWTService) ValidateToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, j.keyFunc)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// keyFunc 用于从JWTService中获取秘钥
func (j *JWTService) keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return j.SecretKey, nil
}
