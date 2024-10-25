package service

import (
	"back-end/config"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(openID string) (string, error) {
	// 设置过期时间，这里设置为 1 小时
	// 创建一个新的 token 对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"openid": openID,
	})
	// 使用密钥进行签名
	tokenString, err := token.SignedString([]byte(config.Jwtkey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Jwtkey), nil
	})
	if err != nil {
		return "", err
	}
	if openid, ok := token.Claims.(jwt.MapClaims)["openid"]; ok && token.Valid {
		return openid.(string), nil
	} else {
		return "", fmt.Errorf("invalid token")
	}
}
