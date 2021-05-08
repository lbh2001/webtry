package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
	"webtry/model"
)

/**
 * @Author: lbh
 * @Date: 2021/5/8
 * @Description:
 */

const (
	tokenExpiredDuration = time.Hour * 24
)

var signature = []byte("lbh!Q@W#E$R%T")

func GetToken(phone string) (string, error) {
	claims := model.Claims{
		phone,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpiredDuration).Unix(),
			Issuer:    "projectHandler",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signature)
}

func ParseToken(tokenString string) (*model.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return signature, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token is expired")
}
