package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/shijiu-xf/go-base/comm/jwtsj"
	"go.uber.org/zap"
	"net/http"
)

func JwtAuthHandlerFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("sj-token")
		parseTk, err := jwtsj.ParseJwtByGlKey(token)
		if !parseTk.Valid {
			validationError, ok := err.(*jwt.ValidationError)
			if !ok {
				zap.L().Error("[JwtTokenAuth] token 解析出错", zap.Error(err))
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部异常",
				})
				ctx.Abort()
				return
			}
			errMsg := ""
			switch {
			case validationError.Errors&jwt.ValidationErrorMalformed != 0:
				errMsg += "令牌格式不正确;"
			case validationError.Errors&jwt.ValidationErrorUnverifiable != 0:
				errMsg += "签名问题，无法验证令牌;"
			case validationError.Errors&jwt.ValidationErrorSignatureInvalid != 0:
				errMsg += "无效的令牌;"
			case validationError.Errors&jwt.ValidationErrorExpired != 0:
				errMsg += "登录已过期;"
			default:
				errMsg = validationError.Error()
			}
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg": "未登录或登录已过期",
			})
			zap.L().Error("[JwtTokenAuth] token 无效", zap.String("tokenErr", errMsg))
			ctx.Abort()
			return
		}
	}
}
