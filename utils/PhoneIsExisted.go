package utils

import (
	"github.com/jinzhu/gorm"
	"webtry/model"
)

/**
 * @Author: lbh
 * @Date: 2021/4/9
 * @Description: 判断手机号是否存在
 */

func PhoneIsExisted(phone string,db *gorm.DB) bool {
	var user model.User
	db.Where("phone = ?",phone).Find(&user)
	if user.ID == 0 {
		return false
	}
	return true
}