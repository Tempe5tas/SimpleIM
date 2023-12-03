package la_rsp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	SUCCESS = 0
	ERROR   = -1

	ParameterValidationError = 1
	UserEmailExist           = 2
	UserEmailRegistering     = 3
	ProcessError             = 4
	VerifyCodeError          = 5
	UserNotExist             = 6
	UserPasswordError        = 7
	VerifyPassword           = 8
	AuthError                = 9
	GroupNotExist            = 10
	PasswordError            = 11
)

var msgMap = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",

	ParameterValidationError: "参数解析错误",
	UserEmailExist:           "邮箱已被注册",
	UserEmailRegistering:     "邮箱在被注册ing",
	ProcessError:             "流程错误",
	VerifyCodeError:          "验证错误",
	UserNotExist:             "用户不存在",
	UserPasswordError:        "密码错误",
	VerifyPassword:           "密码校验错误",
	AuthError:                "权限验证错误",
	GroupNotExist:            "群不存在",
	PasswordError:            "密码错误",
}

// HttpResponse  setting gin.JSON
func HttpResponse(ctx *gin.Context, httpCode, errCode int, data interface{}) {
	switch data.(type) {
	case error:
		ctx.JSON(httpCode, Response{
			Code: errCode,
			Msg:  getMsg(errCode),
			Data: fmt.Sprint(data),
		})
	default:
		ctx.JSON(httpCode, Response{
			Code: errCode,
			Msg:  getMsg(errCode),
			Data: data,
		})
	}
}

func Success(ctx *gin.Context, data interface{}) {
	HttpResponse(ctx, http.StatusOK, 0, data)
	ctx.Abort()
}

func Failed(ctx *gin.Context, errCode int, err error) {
	HttpResponse(ctx, http.StatusOK, errCode, err)
	ctx.Abort()
}

func FailedWithData(ctx *gin.Context, errCode int, data interface{}) {
	HttpResponse(ctx, http.StatusOK, errCode, data)
	ctx.Abort()
}

func getMsg(code int) string {
	msg, ok := msgMap[code]
	if ok {
		return msg
	}
	return msgMap[ERROR]
}
