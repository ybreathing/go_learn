package main

import (
	"zzy/go-learn/controller"
	"zzy/go-learn/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.POST("/user/register", controller.Register)
	r.POST("/user/login", controller.Login)
	r.GET("/user/info", middleware.AuthMiddleware(), controller.UserInfo)
	return r
}
