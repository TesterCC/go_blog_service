package errcode

import (
	"fmt"
	"net/http"
)

// 编写常用的一些错误处理公共方法，标准化项目的错误输出

// 声明了 Error 结构体用于表示错误的响应结果

type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

// 用 codes 作为全局错误码的存储载体，便于查看当前注册情况

var codes = map[int]string{}

// 在调用 NewError 创建新的 Error 实例的同时进行排重的校验。

func NewError(code int, msg string) *Error {
	// 这种语法判断key是否在map中,即code是否在codes map中
	if _, ok := codes[code]; ok {
		// 手动调用内置的 panic 函数去触发 panic 异常，panic 函数可接受任何值作为参数
		// FIXME：update to goland2021.3 or higher 支持完整的go泛型, 这里Sprintf语法报错（编译能过）推测是因为Golang 2021.2.2对go1.18的泛型支持不充分的原因。
		panic(fmt.Sprintf("错误码 %d 已存在，请更换一个使用", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息：%s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}

	return &newError
}

// 主要用于针对一些特定错误码进行状态码的转换，因为不同的内部错误码在 HTTP 状态码中都代表着不同的意义，我们需要将其区分开来，便于客户端以及监控/报警等系统的识别和监听。

func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	case NotFound.Code():
		return http.StatusNoContent
	}

	return http.StatusInternalServerError
}
