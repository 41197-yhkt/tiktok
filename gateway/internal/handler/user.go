package handler

import (
	"context"
	"errors"
	"log"
	"time"

	// "github.com/41197-yhkt/tiktok/gateway/kitex_gen/relation"
	// "github.com/41197-yhkt/tiktok/gateway/kitex_gen/relation/relationservice"
	"github.com/41197-yhkt/tiktok/gateway/biz/model/douyin"
	"github.com/41197-yhkt/tiktok/gateway/kitex_gen/user"
	"github.com/41197-yhkt/tiktok/gateway/kitex_gen/user/userservice"
	"github.com/41197-yhkt/tiktok/pkg/errno"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// DouyinUserRegisterMethod .
// @router /relation/user/register [POST]
func DouyinUserRegisterMethod(ctx context.Context, c *app.RequestContext) {
	//return
	hlog.Info("in UserRegisterMethod")
	var err error
	var reqHTTP douyin.DouyinUserRegisterRequest
	err = c.BindAndValidate(&reqHTTP)
	hlog.Info("user:", reqHTTP)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		SendResponse(c, *errno.ServerError)
		return
	}

	client, err := userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		hlog.Info("err:", err.Error())
		SendResponse(c, *errno.ServerError)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	reqRPC := user.UserRegisterRequest{
		Username: reqHTTP.Username,
		Password: reqHTTP.Password,
	}
	resp, err := client.UserRegister(ctx, &reqRPC)
	cancel()
	if err != nil {
		hlog.Info("err:", err.Error())
		SendResponseWithErr(c, 1, err.Error())
		c.Abort()
		return
	}
	if resp.BaseResp.StatusCode != 0 {
		SendResponseWithErr(c, resp.BaseResp.StatusCode, *resp.BaseResp.StatusMsg)
		return
	}

	// RPC和HTTP的返回一致，就没有更改
	// c.JSON(consts.StatusOK, resp)

}

// DouyinUserLoginMethod .
// @router /relation/user/login [POST]
func DouyinUserLoginMethod(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	// return &user.UserLoginResponse{
	// 	BaseResp: nil,
	// 	UserId:   2,
	// }, nil
	hlog.Info("in user login")
	var err error
	var reqHTTP douyin.DouyinUserRegisterRequest
	err = c.BindAndValidate(&reqHTTP)
	hlog.Info("user:", reqHTTP)
	if err != nil {
		return nil, err
	}

	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		return nil, err
	}

	client, err := userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	reqRPC := user.UserLoginRequest{
		Username: reqHTTP.Username,
		Password: reqHTTP.Password,
	}
	resp, err := client.UserLogin(ctx, &reqRPC)
	cancel()
	if err != nil {
		hlog.Info("err:", err.Error())
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		hlog.Info("BaseResp.StatusMsg:", resp.BaseResp.StatusMsg)
		return nil, errors.New(*resp.BaseResp.StatusMsg)
	}

	hlog.Info("resp:", resp)

	return resp, nil
}

// 用于测试
func DouyinUserLoginMethodTest(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	// return &user.UserLoginResponse{
	// 	BaseResp: nil,
	// 	UserId:   2,
	// }, nil
	var err error
	var req user.UserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		return nil, err
	}
	// r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	// if err != nil {
	// 	return nil, err
	// }
	// client, err := userservice.NewClient("user", client.WithResolver(r))
	// if err != nil {
	// 	return nil, err
	// }
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	// resp, err := client.UserLogin(ctx, &req)
	// cancel()
	// if err != nil {
	// 	return nil, err
	// }
	msg := "用户已存在"
	resp := user.UserLoginResponse{
		BaseResp: &user.BaseResp{
			StatusCode: 11,
			StatusMsg:  &msg,
		},
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errors.New(*resp.BaseResp.StatusMsg)
	}

	return resp, nil
}

// DouyinUserMethod .
// @router /relation/user [GET]
func DouyinUserMethod(ctx context.Context, c *app.RequestContext) {
	hlog.Info("in user Method")
	var err error
	var reqHTTP douyin.DouyinUserRequest
	err = c.BindAndValidate(&reqHTTP)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	hlog.Info("user:", reqHTTP)
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		hlog.Info("err:", err.Error())
		SendResponse(c, *errno.ServerError)
		return
	}

	client, err := userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		hlog.Info("err:", err.Error())
		SendResponse(c, *errno.ServerError)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	reqRPC := user.UserInfoRequest{
		UserId: reqHTTP.UserID,
	}
	resp, err := client.UserInfo(ctx, &reqRPC)
	cancel()
	if err != nil {
		hlog.Info("err:", err.Error())
		SendResponseWithErr(c, 1, err.Error())
		return
	}
	if resp.BaseResp.StatusCode != 0 {
		SendResponseWithErr(c, resp.BaseResp.StatusCode, *resp.BaseResp.StatusMsg)
		return
	}
	//RPC和HTTP返回的一致，没有更改
	c.JSON(consts.StatusOK, resp)
}

// DouyinRelationActionMethod .
// @router /relation/relation/action [POST]
// relation移动到了user中
func DouyinRelationActionMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var reqHTTP douyin.DouyinRelationActionRequest
	err = c.BindAndValidate(&reqHTTP)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	user_id := getUserIdFromJWT(ctx, c)
	hlog.Info("to_user_id=", reqHTTP.ToUserID, " user_id=", user_id)

	reqRPC := user.RelationActionRequest{
		UserId:     user_id,
		ToUserId:   reqHTTP.ToUserID,
		ActionType: reqHTTP.ActionType,
	}

	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		SendResponse(c, *errno.ServerError)
		return
	}
	client, err := userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		SendResponse(c, *errno.ServerError)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	resp, err := client.RealtionAction(ctx, &reqRPC)
	cancel()
	if err != nil {
		SendResponse(c, *errno.ServerError)
		return
	}
	if resp.BaseResp.StatusCode != 0 {
		SendResponseWithErr(c, resp.BaseResp.StatusCode, *resp.BaseResp.StatusMsg)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// DouyinRelationFollowListMethod .
// @router /relation/relation/follow/list [GET]
func DouyinRelationFollowListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.FollowListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		SendResponse(c, *errno.ServerError)
		return
	}
	client, err := userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		SendResponse(c, *errno.ServerError)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	resp, err := client.GetFollowList(ctx, &req)
	cancel()
	if err != nil {
		SendResponse(c, *errno.ServerError)
		return
	}
	if resp.BaseResp.StatusCode != 0 {
		SendResponseWithErr(c, resp.BaseResp.StatusCode, *resp.BaseResp.StatusMsg)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// DouyinRelationFollowerListMethod .
// @router /relation/relation/follower/list [GET]
func DouyinRelationFollowerListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.FollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		SendResponse(c, *errno.ServerError)
		return
	}
	client, err := userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		SendResponse(c, *errno.ServerError)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	resp, err := client.GetFollowerList(ctx, &req)
	cancel()
	if err != nil {
		SendResponse(c, *errno.ServerError)
		return
	}
	if resp.BaseResp.StatusCode != 0 {
		SendResponseWithErr(c, resp.BaseResp.StatusCode, *resp.BaseResp.StatusMsg)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// DouyinRelationFriendListMethod .
// @router /relation/relation/friend/list [GET]
func DouyinRelationFriendListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.FriendListRequest
	err = c.BindAndValidate(&req)
	log.Print(req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		SendResponse(c, *errno.ServerError)
		return
	}
	client, err := userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		SendResponse(c, *errno.ServerError)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	resp, err := client.GetFriendList(ctx, &req)
	cancel()
	if err != nil {
		SendResponse(c, *errno.ServerError)
		return
	}
	if resp.BaseResp.StatusCode != 0 {
		SendResponseWithErr(c, resp.BaseResp.StatusCode, *resp.BaseResp.StatusMsg)
		return
	}
	c.JSON(consts.StatusOK, resp)
}
