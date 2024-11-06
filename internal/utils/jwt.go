package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Uuid    string `json:"uuid"`
	IsAdmin bool   `json:"isAdmin"`
	jwt.StandardClaims
}

type TokenHandler struct {
	JwtKey []byte
}

// NewTokenHandler 返回一个TokenHandler
func NewTokenHandler() TokenHandler {
	return TokenHandler{}
}

func (th TokenHandler) GenerateToken(uuid string, isAdmin bool) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // 设置过期时间
	claims := &Claims{
		Uuid:    uuid,
		IsAdmin: isAdmin,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(th.JwtKey)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}

func (th TokenHandler) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString,
		&Claims{},
		func(t *jwt.Token) (interface{}, error) {
			return th.JwtKey, nil
		})
	if err != nil {
		return nil, err
	}
	if Claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return Claims, nil
	}
	return nil, errors.New("未知的token")
}
