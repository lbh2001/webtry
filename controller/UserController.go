package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"webtry/common"
	"webtry/model"
	"webtry/utils"
)

/**
 * @Author: lbh
 * @Date: 2021/4/9
 * @Description:
 */

//注册
func Register(c *gin.Context) {

	db := common.GetDB()

	//获取参数
	username := c.PostForm("username")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	//手机号必须为11位
	if len(phone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "fail",
			"msg":    "手机号必须为11位",
		})
		return
	}

	//密码至少为6位
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "fail",
			"msg":    "密码不能少于6位！",
		})
		return
	}

	//若名称为空，则给一个10位的随机字符串
	if len(username) == 0 {
		username = utils.RandomString(10)
	}

	//判断手机号是否存在
	if utils.PhoneIsExisted(phone, db) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "fail",
			"msg":    "该手机号已被注册！",
		})
		return
	}

	//密码加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"msg":    "密码加密出错！",
		})
		return
	}

	//创建用户
	var user = model.User{
		Username: username,
		Password: string(hasedPassword),
		Phone:    phone,
	}

	db.Debug().Create(&user)

	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"message": "success to register!",
	})
}

func Login(c *gin.Context) {
	//获取参数
	db := common.GetDB()

	password := c.PostForm("password")
	phone := c.PostForm("phone")

	//手机号必须为11位
	if len(phone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "fail",
			"msg":    "手机号必须为11位",
		})
		return
	}

	//密码至少为6位
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "fail",
			"msg":    "密码不能少于6位！",
		})
		return
	}

	//根据手机号查询用户
	var user model.User
	db.Where("phone = ?", phone).Find(&user)

	//若用户不存在
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "fail",
			"msg":    "用户不存在！",
		})
		return
	}

	//解密、判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"msg":    "密码错误！",
		})
		return
	}

	//发放token
	token, _ := utils.GetToken(phone)
	fmt.Println(token)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"msg":    "登录成功!",
		"data":   gin.H{"token": token},
	})

}

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success!",
	})
}
