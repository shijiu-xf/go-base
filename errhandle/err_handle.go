package errhandle

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	enT "github.com/go-playground/locales/en"
	zhT "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
	"github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

// CustomErrors 自定义错误结构体
// FieldName 字段名(结构体名.字段名 例如 EmailLogin.Password)
// Tag validator 组件提供的标签名 例如required,max
// CustomMsg 输出的错误信息
type CustomErrors struct {
	// FieldName 校验的字段名
	FieldName string
	// Validation Tag 的名字
	Tag string
	// CustomMsg 自定义的返回信息
	CustomMsg string
}

func GrpcErrToHttpErr(ctx *gin.Context, err error) {
	if err == nil {
		return
	}
	zap.L().Error("请求返回失败", zap.String("msg", err.Error()))
	if e, ok := status.FromError(err); ok {
		switch e.Code() {

		case codes.NotFound:
			ctx.JSON(http.StatusNotFound, gin.H{
				"msg": e.Message(),
			})
		case codes.InvalidArgument:
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": e.Message(),
			})
		case codes.AlreadyExists:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": e.Message(),
			})
		case codes.Unimplemented:
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg": "无权操作",
			})
		case codes.Internal:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "系统内部异常",
			})
		case codes.Unavailable:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "服务不可用",
			})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "其他异常",
			})
		}
	}
}

// ValidationCustomErrors 自定义字段校验失败的返回信息
func ValidationCustomErrors(ctx *gin.Context, err error, st ...[]CustomErrors) {
	fieldErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "内部错误",
		})
		zap.L().Error("[ValidationCustomErrors] 转换错误", zap.String("原错误信息", err.Error()))
		return
	}
	zap.L().Error("[字段校验] 原错误信息", zap.String("msg", err.Error()))
	// 转map便于匹配 ,字段名+tag名作为key
	fieldErrMap := fieldErrsTOMap(fieldErrs)
	customMsg := ""
	if len(st) > 0 {
		for _, cuErr := range st[0] {
			if fieldErrMap[cuErr.FieldName+cuErr.Tag] != nil {
				customMsg = customMsg + cuErr.CustomMsg + ";"
				// 匹配上了就删除，方便后续循环出没有自定义的错误信息返回
				delete(fieldErrMap, cuErr.FieldName+cuErr.Tag)
			}
		}
	}
	// 没有设置自定义错误的错误信息拼接
	for _, fieldError := range fieldErrMap {
		customMsg = customMsg + fieldError.Translate(GetTrans("zh")) + ";"
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"msg": customMsg,
	})
}

// 将 ValidationErrors 转为map Field+Tag 作为key
func fieldErrsTOMap(fieldErrs []validator.FieldError) map[string]validator.FieldError {
	m := make(map[string]validator.FieldError, 0)
	for _, fieldErr := range fieldErrs {
		if m[fieldErr.Namespace()+fieldErr.Tag()] == nil {
			m[fieldErr.Namespace()+fieldErr.Tag()] = fieldErr
		}
	}
	return m
}

func GetTrans(locale string) ut.Translator {
	validate, _ := binding.Validator.Engine().(*validator.Validate)
	ZH := zhT.New()           // 中文翻译器
	EN := enT.New()           // 英文翻译器
	UNI := ut.New(EN, ZH, EN) // 第一个为备用语言环境
	trans, ok := UNI.GetTranslator(locale)
	deftT, _ := UNI.GetTranslator("en")
	if !ok {
		zap.L().Error("[InitTrans] 不存在的翻译器,将使用默认的英文翻译器", zap.String("不存在的翻译器为", locale))
		trans = deftT
	}
	var err error
	switch locale {
	case "zh":
		err = zh.RegisterDefaultTranslations(validate, trans)
	case "en":
		err = en.RegisterDefaultTranslations(validate, trans)
	default:
		err = en.RegisterDefaultTranslations(validate, trans)
	}
	if err != nil {
		zap.L().Error("[InitTrans] 翻译器注册失败", zap.Error(err))
	}
	return trans

}
