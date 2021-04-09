package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"webtry/model"
)

/**
 * @Author: lbh
 * @Date: 2021/4/9
 * @Description:数据库初始化
 */

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	username := "root"
	password := "1178055813"
	host := "localhost"
	port := "3306"
	database := "webtry"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(driverName,args)
	if err != nil {
		panic("fail to connect database,err" + err.Error())
	}

	db.SingularTable(true)

	db.AutoMigrate(&model.User{})

	DB = db

	return db
}

func GetDB() *gorm.DB{
	return DB
}