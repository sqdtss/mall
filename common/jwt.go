package common

import (
	"github.com/golang-jwt/jwt"
	"mall/global"
	"time"
)

var SigningKey = []byte(global.Config.Jwt.SigningKey)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken 生成Token
func GenerateToken(username string) (string, error) {
	claims := Claims{username, jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + 60*60,
		Issuer:    username,
	},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SigningKey)
}

// VerifyToken 验证Token
func VerifyToken(tokenString string) error {
	_, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})
	return err
}
