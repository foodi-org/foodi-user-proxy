package response

import (
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type (
	Body struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	CodeErr struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

// Error
//
//	@Description: 实现 error 接口
func (c *CodeErr) Error() string {
	return c.Message
}

// NewCodeErr
//
//	@Description: 定义抛出的异常,支持调用传递
//	@param code 错误码
//	@param message 返回 message
//	@return error
func NewCodeErr(code int, message string) error {
	return &CodeErr{
		Code:    code,
		Message: message,
	}
}

// New
//
//	@Description: 提供 new 方法，任意地方传递参数返回 CodeError 类型的数据
//	@param code
//	@param message
//	@return CodeErr
func New(code int, message string) CodeErr {
	return CodeErr{
		Code:    code,
		Message: message,
	}
}

// Response
//
//	@Description: 统一返回入口
//	@param w ResponseWriter 对象
//	@param data 返回数据
//	@param err 返回错误
func Response(w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		httpx.OkJson(w, ErrHandler(err))
	} else {
		httpx.OkJson(w, Body{
			Code:    SuccessResponse.Code,
			Message: SuccessResponse.Message,
			Data:    data,
		})
	}
}

// BindFailResponse
//
//	@Description: 绑定参数失败时的返回
//	@param err
func BindFailResponse(w http.ResponseWriter, err error) {
	if err != nil {
		httpx.OkJson(w, Body{Code: BindParamError.Code, Message: err.Error()})
	} else {
		httpx.OkJson(w, Body{Code: BindParamError.Code, Message: BindParamError.Message})
	}
}

// ErrorResponse 返回给前端的数据
func (c *CodeErr) ErrorResponse() CodeErr {
	return CodeErr{
		Code:    c.Code,
		Message: c.Message,
	}
}

// DefaultErrHandler
//
//	@Description: 默认异常状态码函数, 只需传递错误message即可，默认返回code 500
//	@param message
//	@return error
func DefaultErrHandler(message string) error {
	return &CodeErr{
		Code:    ServerError.Code,
		Message: message,
	}
}

// ErrHandler
//
//	@Description: 自定义错误返回函数,错误函数主入口
//	@param err
//	@return interface{}
func ErrHandler(err error) interface{} {
	var codeErr *CodeErr
	switch {
	case errors.As(err, &codeErr):
		return err
	default:
		return CodeErr{Code: ServerError.Code, Message: err.Error()}
	}
}
