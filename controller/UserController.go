package controller

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"zzy/go-learn/common"
	"zzy/go-learn/dto"
	"zzy/go-learn/module"
	"zzy/go-learn/response"
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
		//cxt.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		response.Response(cxt, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}

	if len(password) < 6 {
		//cxt.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码长度不能小于6位"})
		response.Response(cxt, http.StatusUnprocessableEntity, 422, nil, "密码长度不能小于6位")

		return
	}

	if len(name) == 0 {
		name = util.RandomString(10)
	}

	if isPhoneExist(DB, phone) {
		//cxt.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在"})
		response.Response(cxt, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}

	// 对密码进行加密存储
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		//cxt.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "密码加密失败"})
		response.Response(cxt, http.StatusUnprocessableEntity, 500, nil, "密码加密失败")
		return
	}
	newUser := module.User{
		Name:     name,
		Phone:    phone,
		Password: string(fromPassword),
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
		response.Response(cxt, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}

	if len(password) < 6 {
		response.Response(cxt, http.StatusUnprocessableEntity, 422, nil, "密码长度不能小于6位")
		return
	}
	// 校验用户
	var user module.User
	DB.Where("phone=?", phone).First(&user)
	if user.ID == 0 {
		//cxt.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		response.Response(cxt, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
	}
	// 校验密码
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		//cxt.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		response.Fail(cxt, nil, "密码错误")
		return
	}
	// 返回token
	token, err := common.ReleaseToken(user)
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"code": 400, "msg": "生成token失败"})
		log.Printf("token generate error : %v", err)
		return
	}

	// 返回结果
	response.Success(cxt, gin.H{"code": 200, "msg": "登录成功", "token": token}, "")

}

func UserInfo(cxt *gin.Context) {

	// 获取请求头中的用户信息
	user, _ := cxt.Get("user")

	// 返回结果
	//cxt.JSON(http.StatusOK, gin.H{
	//	"code": 200,
	//	//user.(module.User)
	//	"user": dto.ToUserDto(user.(module.User)),
	//})
	response.Success(cxt, gin.H{"code": 200, "user": dto.ToUserDto(user.(module.User))}, "")

}

func isPhoneExist(db *gorm.DB, phone string) bool {
	var user module.User
	db.Where("phone=?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false

}
