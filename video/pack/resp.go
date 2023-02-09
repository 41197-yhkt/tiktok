package pack

import (
	"errors"

	"github.com/41197-yhkt/tiktok-video/kitex_gen/video"

	errno "github.com/41197-yhkt/pkg/errno"
)

func BuildBaseResp(err error) *video.BaseResp {
	if err == nil {
		return baseResp(*errno.Success)
	}

	e := &errno.ErrNo{}

	if errors.As(err, &e) {
		return baseResp(*e)
	}

	return baseResp(*errno.ServerError)
}

func baseResp(errNo errno.ErrNo) *video.BaseResp {
	return &video.BaseResp{StatusCode: int32(errNo.Code), StatusMsg: &errNo.Msg}
}
