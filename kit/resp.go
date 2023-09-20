package kit

import (
	"github.com/nwuw/dev/err"
	"net/http"
)

type Response struct {
	Code    int         `json:"code" xml:"code" yaml:"code"`
	Message string      `json:"message" xml:"message" yaml:"message"`
	Data    interface{} `json:"data" xml:"data" yaml:"data"`
}

type Context interface {
	JSON(code int, obj interface{})
	AbortWithStatusJSON(code int, jsonObj interface{})
}

func SuccessMsg(data interface{}) Response {
	result := Response{
		Code:    err.SUCCESS.GetCode(),
		Message: err.SUCCESS.GetMessage(),
		Data:    data,
	}
	return result
}
func ErrorMsg(err err.Error) Response {
	result := Response{
		Code:    err.GetCode(),
		Message: err.GetMessage(),
		Data:    nil,
	}
	return result
}

func Ok(ctx Context, data interface{}) {
	ctx.JSON(http.StatusOK, SuccessMsg(data))
}

func Err(ctx Context, err err.Error) {
	ctx.AbortWithStatusJSON(http.StatusOK, ErrorMsg(err))
}
