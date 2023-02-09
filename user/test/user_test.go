package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/41197-yhkt/tiktok-user/kitex_gen/user"
	"github.com/41197-yhkt/tiktok-user/service"
)

func TestRegister(t *testing.T) {
	ctx := context.Background()
	req := user.NewUserRegisterRequest()
	req.Username = "ljf_test"
	req.Password = "123456"
	resp, err := service.UserRegister(ctx, req)
	fmt.Println("TestRegister resp = ", resp, " err = ", err)
}

func TestReregister(t *testing.T) {
	ctx := context.Background()
	req := user.NewUserRegisterRequest()
	req.Username = "lyh_test"
	req.Password = "123456"
	resp, err := service.UserRegister(ctx, req)
	fmt.Println("TestRegister resp = ", resp, " err = ", err)
}

func TestLoginSuccess(t *testing.T) {
	ctx := context.Background()
	req := user.NewUserLoginRequest()
	req.Username = "lyh_test"
	req.Password = "123456"
	resp, err := service.UserLogin(ctx, req)
	fmt.Println("TestLoginSuccess resp = ", resp, " err = ", err)
}

func TestLoginFailedWrongPassword(t *testing.T) {
	ctx := context.Background()
	req := user.NewUserLoginRequest()
	req.Username = "lyh_test"
	req.Password = "12345678"
	resp, err := service.UserLogin(ctx, req)
	fmt.Println("TestLoginFailed resp = ", resp, " err = ", err)
}

func TestLoginFailedWrongUsername(t *testing.T) {
	ctx := context.Background()
	req := user.NewUserLoginRequest()
	req.Username = "lyh_test_2"
	req.Password = "123456"
	resp, err := service.UserLogin(ctx, req)
	fmt.Println("TestLoginFailed resp = ", resp, " err = ", err)
}

func TestUserInfo(t *testing.T) {
	ctx := context.Background()
	req := user.NewUserLoginRequest()
	req.Username = "lyh_test"
	req.Password = "123456"
	resp, err := service.UserLogin(ctx, req)
	fmt.Println("TestLoginSuccess resp = ", resp, " err = ", err)

	existenceUserId := resp.UserId
	infoReq := user.NewUserInfoRequest()
	infoReq.UserId = existenceUserId
	userInfo, err := service.UserInfo(ctx, infoReq)
	fmt.Println("TestUserInfo resp = ", userInfo, " err = ", err)
}

func TestGetFollowList(t *testing.T) {
	ctx := context.Background()
	req := user.NewFollowListRequest()
	req.UserId = 1
	resp, err := service.GetFollowList(ctx, req)
	fmt.Println("TestGetFollowList resp = ", resp, " err = ", err)
}

func TestGetFollowerList(t *testing.T) {
	ctx := context.Background()
	req := user.NewFollowerListRequest()
	req.UserId = 1
	resp, err := service.GetFollowerList(ctx, req)
	fmt.Println("TestGetFollowerList resp = ", resp, " err = ", err)
}

func TestGetFriendList(t *testing.T) {
	ctx := context.Background()
	req := user.NewFriendListRequest()
	req.UserId = 1
	resp, err := service.GetFriendList(ctx, req)
	fmt.Println("TestGetFriendList resp = ", resp, " err = ", err)
}

func TestIsFriend(t *testing.T) {
	ctx := context.Background()
	UserId := 1
	ToUserId := 2
	resp, err := service.IsFriend(ctx, int64(UserId), int64(ToUserId))
	fmt.Println("TestGetFriendList resp = ", resp, " err = ", err)
}
