package validator

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"regexp"
)

// ValidateMobile 自定义的国内手机号格式验证器
func ValidateMobile(level validator.FieldLevel) bool {
	mobile := level.Field().String()
	ok, err := regexp.MatchString(`0?(13|14|15|18|17)[0-9]{9}`, mobile)
	if err != nil {
		zap.L().Error("[ValidateMobile] 验证器，正则匹配出错", zap.String("msg", err.Error()))
	}
	return ok
}
