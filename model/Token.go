package model

/**
 * @Author: lbh
 * @Date: 2021/5/8
 * @Description:
 */
import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Phone              string `json:"phone"`
	jwt.StandardClaims `json:"claims"`
}
