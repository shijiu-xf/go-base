package uniform

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinUniformSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "操作成功",
		"data": data,
	})
}

func GinUniformErr(ctx *gin.Context, errCode int, errMsg string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"code": errCode,
		"msg":  errMsg,
		"data": nil,
	})
}
