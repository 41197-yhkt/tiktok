package handler

import (
	"net/http"
	"github.com/41197-yhkt/pkg/errno"
	douyin "tiktok-gateway/internal/model"

	"github.com/cloudwego/hertz/pkg/app"
	//"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// SendResponse pack response
func SendResponse(c *app.RequestContext, err errno.ErrNo) {
	c.JSON(http.StatusOK, douyin.BaseResp{
		StatusCode:    int32(err.Code),
		StatusMsg: &err.Msg,
	})
}
