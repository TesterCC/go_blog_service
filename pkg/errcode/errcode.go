package errcode

import "fmt"

// 编写常用的一些错误处理公共方法，标准化项目的错误输出
type Error struct {
	code int `json:"code"`
	msg string `json:"msg"`
	details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	// 判断key是否在map中,即code是否在codes map中
	if _,ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码%d已存在，请更好一个使用", code))
	}
// https://github.com/go-programming-tour-book/blog-service/blob/master/pkg/errcode/errcode.go
// https://golang2.eddycjy.com/posts/ch2/03-auxiliary-component/
}