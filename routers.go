package main

import (
	"zzy/go-learn/controller"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.POST("/user/register", controller.Register)
	r.POST("/user/login", controller.Login)
	return r
}
