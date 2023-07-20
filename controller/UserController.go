package controller

import (
	"net/http"
	"zzy/go-learn/common"
	"zzy/go-learn/module"
	"zzy/go-learn/util"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Register(cxt *gin.Context) {

	DB := common.GetDB()
	// 获取接口去请求参数
	name := cxt.PostForm("name")
	phone := cxt.PostForm("phone")
	password := cxt.PostForm("password")

	// 判空
	if len(phone) != 11 {
		//cxt.JSON(200, map[string]interface{}{"code": 422, "msg": ""})
		cxt.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}

	if len(password) < 6 {
		cxt.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码长度不能小于6位"})
		return
	}

	if len(name) == 0 {
		name = util.RandomString(10)
		return
	}

	if isPhoneExist(DB, phone) {
		cxt.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在"})
		return
	}

	newUser := module.User{
		Name:     name,
		Phone:    phone,
		Password: password,
	}
	DB.Create(&newUser)

	cxt.JSON(200, gin.H{
		"msg": "注册成功",
	})

}

func Login(cxt *gin.Context) {
	// 获取DB
	DB := common.GetDB()

	//获取参数
	// 获取接口去请求参数
	phone := cxt.PostForm("phone")
	password := cxt.PostForm("password")

	// 校验参数
	// 判空
	if len(phone) != 11 {
		//cxt.JSON(200, map[string]interface{}{"code": 422, "msg": ""})
		cxt.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}

	if len(password) < 6 {
		cxt.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码长度不能小于6位"})
		return
	}
	// 校验用户
	var user module.User
	DB.Where("phone=?", phone).First(&user)
	if user.ID == 0 {
		cxt.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
	}

	// 校验密码

	// 返回token

	// 返回结果

	//

	cxt.JSON(200, gin.H{
		"message": "pong",
	})

}

func isPhoneExist(db *gorm.DB, phone string) bool {
	var user module.User
	db.Where("phone=?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false

}
