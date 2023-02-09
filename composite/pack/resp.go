package pack

import (
	"errors"

	"github.com/41197-yhkt/tiktok/tiktok-composite/kitex_gen/composite"

	"github.com/41197-yhkt/tiktok/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *composite.BaseResp {
	if err == nil {
		return baseResp(*errno.Success)
	}

	e := &errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(*e)
	}

	return baseResp(*errno.ServerError)
}

func baseResp(err errno.ErrNo) *composite.BaseResp {
	return &composite.BaseResp{StatusCode: int32(err.Code), StatusMsg: &err.Msg}
}
