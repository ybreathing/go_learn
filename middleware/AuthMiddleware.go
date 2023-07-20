package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"zzy/go-learn/common"
	"zzy/go-learn/module"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取请求头中的token
		tokenString := context.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "请求头token异常"})
			context.Abort()
			return
		}
		tokenString = tokenString[7:]

		token, claims, err := common.ParesToken(tokenString)
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "token校验失败"})
			context.Abort()
			return
		}
		userId := claims.UserId
		DB := common.GetDB()
		var user module.User
		DB.First(&user, userId)
		if user.ID == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "用户信息不存在"})
			context.Abort()
			return
		}

		// 将用户信息放请求头中
		context.Set("user", user)
		context.Next()

	}
}
