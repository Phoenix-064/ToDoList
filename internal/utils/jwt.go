package utils

import (
	"errors"
	"strings"
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

// GenerateToken 生成一个token
func (th TokenHandler) GenerateToken(uuid string, isAdmin bool) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour).Unix() // 设置过期时间
	claims := &Claims{
		Uuid:    uuid,
		IsAdmin: isAdmin,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expirationTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(th.JwtKey)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}

// parseToken 解析token
func (th TokenHandler) parseToken(tokenString string) (*Claims, error) {
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

// ValidateToken 验证token是否可用，同时返回解析后的结果
func (th TokenHandler) ValidateToken(tokenString string) (*Claims, error) {
	if tokenString == "" {
		return nil, errors.New("空的token")
	}
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, errors.New("token格式错误")
	}
	c, err := th.parseToken(tokenString)
	if err != nil {
		return nil, err
	}
	if c.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("已过期的token")
	}
	if c.Uuid == "" {
		return nil, errors.New("空的uuid")
	}
	return &Claims{}, nil
}
