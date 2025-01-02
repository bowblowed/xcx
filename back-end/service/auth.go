package service

import (
	"back-end/config"
	"back-end/model"
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID": id,
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
	m := token.Claims.(jwt.MapClaims)
	if id, ok := m["ID"]; ok && token.Valid {
		return id.(string), nil
	} else {
		return "", fmt.Errorf("invalid token")
	}
}

func Auth(auth string) (model.User, error) {
	ID, err := ParseToken(auth)
	if err != nil {
		return model.User{}, err
	}
	id, err := strconv.Atoi(ID)
	if err != nil {
		return model.User{}, err
	}
	return GetUser(uint(id))
}
