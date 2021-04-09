package utils

import (
	"math/rand"
	"time"
)

/**
 * @Author: lbh
 * @Date: 2021/4/9
 * @Description:
 */

//生成n位随机字符串
func RandomString(n int) string {

	rand.Seed(time.Now().Unix())

	var letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	result := make([]byte,n)
	for i := range result{
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}