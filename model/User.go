package model

import "github.com/jinzhu/gorm"

/**
 * @Author: lbh
 * @Date: 2021/4/9
 * @Description:
 */
type User struct {
	gorm.Model
	Username string	`gorm:"type:varchar(30);not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Phone	 string `gorm:"type:varchar(11);not null;unique"`
}