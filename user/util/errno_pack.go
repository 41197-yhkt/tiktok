package util

import (
	"errors"
	"github.com/41197-yhkt/pkg/errno"
	"github.com/41197-yhkt/tiktok-user/kitex_gen/user"
)

func PackBaseResp(err error) *user.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := &errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	sErr := errno.ServerError
	if err.Error() != "" {
		sErr.Msg = err.Error()
	}
	return baseResp(sErr)
}

func baseResp(err *errno.ErrNo) *user.BaseResp {
	return &user.BaseResp{
		StatusCode: int32(err.Code),
		StatusMsg:  &err.Msg,
	}
}
