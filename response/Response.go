package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(cxt *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	cxt.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(cxt *gin.Context, data gin.H, msg string) {
	Response(cxt, http.StatusOK, 200, data, msg)
}

func Fail(cxt *gin.Context, data gin.H, msg string) {
	Response(cxt, http.StatusOK, 400, data, msg)
}
