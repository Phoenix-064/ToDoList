package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserId string
	jwt.StandardClaims
}

type TokenHandler struct {
	JwtKey []byte
}

func (th TokenHandler) GenerateToken(userId string) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // 设置过期时间
	claims := &Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(th.JwtKey)
}
