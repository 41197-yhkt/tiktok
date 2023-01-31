package errno

import (
	"fmt"
	"net/http"
)

type ErrNo struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

func (e ErrNo)Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息：%s", e.Code, e.Msg) 
}

func NewErrNo(code int, msg string) *ErrNo{
	return &ErrNo{Code:code, Msg: msg}
}

// 转换成http错误码
func (e *ErrNo) StatusCode() int {
	switch e.Code{
	case Success.Code:
		return http.StatusOK
	case ServerError.Code:
		return http.StatusInternalServerError
	case InvalidParams.Code:
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code:
		fallthrough
	case UnauthorizedTokenError.Code:
		fallthrough
	case UnauthorizedTokenGenerate.Code:
		fallthrough
	case UnauthorizedTokenTimeout.Code:
		return http.StatusUnauthorized
	case TooManyRequests.Code:
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}