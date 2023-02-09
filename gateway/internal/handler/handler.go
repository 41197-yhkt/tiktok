package handler

import (
	"net/http"

	douyin "github.com/41197-yhkt/tiktok/tiktok-gateway/internal/model"

	"github.com/41197-yhkt/tiktok/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	//"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// SendResponse pack response
func SendResponse(c *app.RequestContext, err errno.ErrNo) {
	c.JSON(http.StatusOK, douyin.BaseResp{
		StatusCode: int32(err.Code),
		StatusMsg:  &err.Msg,
	})
}
