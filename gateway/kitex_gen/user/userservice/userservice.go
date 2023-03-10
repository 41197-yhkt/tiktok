// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	user "github.com/41197-yhkt/tiktok/gateway/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"UserRegister":    kitex.NewMethodInfo(userRegisterHandler, newUserServiceUserRegisterArgs, newUserServiceUserRegisterResult, false),
		"UserLogin":       kitex.NewMethodInfo(userLoginHandler, newUserServiceUserLoginArgs, newUserServiceUserLoginResult, false),
		"UserInfo":        kitex.NewMethodInfo(userInfoHandler, newUserServiceUserInfoArgs, newUserServiceUserInfoResult, false),
		"UserFollow":      kitex.NewMethodInfo(userFollowHandler, newUserServiceUserFollowArgs, newUserServiceUserFollowResult, false),
		"UserUnfollow":    kitex.NewMethodInfo(userUnfollowHandler, newUserServiceUserUnfollowArgs, newUserServiceUserUnfollowResult, false),
		"GetFollowList":   kitex.NewMethodInfo(getFollowListHandler, newUserServiceGetFollowListArgs, newUserServiceGetFollowListResult, false),
		"GetFollowerList": kitex.NewMethodInfo(getFollowerListHandler, newUserServiceGetFollowerListArgs, newUserServiceGetFollowerListResult, false),
		"GetFriendList":   kitex.NewMethodInfo(getFriendListHandler, newUserServiceGetFriendListArgs, newUserServiceGetFriendListResult, false),
		"IsFriend":        kitex.NewMethodInfo(isFriendHandler, newUserServiceIsFriendArgs, newUserServiceIsFriendResult, false),
		"CompGetUser":     kitex.NewMethodInfo(compGetUserHandler, newUserServiceCompGetUserArgs, newUserServiceCompGetUserResult, false),
		"CompMGetUser":    kitex.NewMethodInfo(compMGetUserHandler, newUserServiceCompMGetUserArgs, newUserServiceCompMGetUserResult, false),
		"RelationAction":  kitex.NewMethodInfo(relationActionHandler, newUserServiceRelationActionArgs, newUserServiceRelationActionResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func userRegisterHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserRegisterArgs)
	realResult := result.(*user.UserServiceUserRegisterResult)
	success, err := handler.(user.UserService).UserRegister(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserRegisterArgs() interface{} {
	return user.NewUserServiceUserRegisterArgs()
}

func newUserServiceUserRegisterResult() interface{} {
	return user.NewUserServiceUserRegisterResult()
}

func userLoginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserLoginArgs)
	realResult := result.(*user.UserServiceUserLoginResult)
	success, err := handler.(user.UserService).UserLogin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserLoginArgs() interface{} {
	return user.NewUserServiceUserLoginArgs()
}

func newUserServiceUserLoginResult() interface{} {
	return user.NewUserServiceUserLoginResult()
}

func userInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserInfoArgs)
	realResult := result.(*user.UserServiceUserInfoResult)
	success, err := handler.(user.UserService).UserInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserInfoArgs() interface{} {
	return user.NewUserServiceUserInfoArgs()
}

func newUserServiceUserInfoResult() interface{} {
	return user.NewUserServiceUserInfoResult()
}

func userFollowHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserFollowArgs)
	realResult := result.(*user.UserServiceUserFollowResult)
	success, err := handler.(user.UserService).UserFollow(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserFollowArgs() interface{} {
	return user.NewUserServiceUserFollowArgs()
}

func newUserServiceUserFollowResult() interface{} {
	return user.NewUserServiceUserFollowResult()
}

func userUnfollowHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserUnfollowArgs)
	realResult := result.(*user.UserServiceUserUnfollowResult)
	success, err := handler.(user.UserService).UserUnfollow(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserUnfollowArgs() interface{} {
	return user.NewUserServiceUserUnfollowArgs()
}

func newUserServiceUserUnfollowResult() interface{} {
	return user.NewUserServiceUserUnfollowResult()
}

func getFollowListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetFollowListArgs)
	realResult := result.(*user.UserServiceGetFollowListResult)
	success, err := handler.(user.UserService).GetFollowList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetFollowListArgs() interface{} {
	return user.NewUserServiceGetFollowListArgs()
}

func newUserServiceGetFollowListResult() interface{} {
	return user.NewUserServiceGetFollowListResult()
}

func getFollowerListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetFollowerListArgs)
	realResult := result.(*user.UserServiceGetFollowerListResult)
	success, err := handler.(user.UserService).GetFollowerList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetFollowerListArgs() interface{} {
	return user.NewUserServiceGetFollowerListArgs()
}

func newUserServiceGetFollowerListResult() interface{} {
	return user.NewUserServiceGetFollowerListResult()
}

func getFriendListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetFriendListArgs)
	realResult := result.(*user.UserServiceGetFriendListResult)
	success, err := handler.(user.UserService).GetFriendList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetFriendListArgs() interface{} {
	return user.NewUserServiceGetFriendListArgs()
}

func newUserServiceGetFriendListResult() interface{} {
	return user.NewUserServiceGetFriendListResult()
}

func isFriendHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceIsFriendArgs)
	realResult := result.(*user.UserServiceIsFriendResult)
	success, err := handler.(user.UserService).IsFriend(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceIsFriendArgs() interface{} {
	return user.NewUserServiceIsFriendArgs()
}

func newUserServiceIsFriendResult() interface{} {
	return user.NewUserServiceIsFriendResult()
}

func compGetUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceCompGetUserArgs)
	realResult := result.(*user.UserServiceCompGetUserResult)
	success, err := handler.(user.UserService).CompGetUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceCompGetUserArgs() interface{} {
	return user.NewUserServiceCompGetUserArgs()
}

func newUserServiceCompGetUserResult() interface{} {
	return user.NewUserServiceCompGetUserResult()
}

func compMGetUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceCompMGetUserArgs)
	realResult := result.(*user.UserServiceCompMGetUserResult)
	success, err := handler.(user.UserService).CompMGetUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceCompMGetUserArgs() interface{} {
	return user.NewUserServiceCompMGetUserArgs()
}

func newUserServiceCompMGetUserResult() interface{} {
	return user.NewUserServiceCompMGetUserResult()
}

func relationActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceRelationActionArgs)
	realResult := result.(*user.UserServiceRelationActionResult)
	success, err := handler.(user.UserService).RelationAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceRelationActionArgs() interface{} {
	return user.NewUserServiceRelationActionArgs()
}

func newUserServiceRelationActionResult() interface{} {
	return user.NewUserServiceRelationActionResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (r *user.UserRegisterResponse, err error) {
	var _args user.UserServiceUserRegisterArgs
	_args.Req = req
	var _result user.UserServiceUserRegisterResult
	if err = p.c.Call(ctx, "UserRegister", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserLogin(ctx context.Context, req *user.UserLoginRequest) (r *user.UserLoginResponse, err error) {
	var _args user.UserServiceUserLoginArgs
	_args.Req = req
	var _result user.UserServiceUserLoginResult
	if err = p.c.Call(ctx, "UserLogin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserInfo(ctx context.Context, req *user.UserInfoRequest) (r *user.UserInfoResponse, err error) {
	var _args user.UserServiceUserInfoArgs
	_args.Req = req
	var _result user.UserServiceUserInfoResult
	if err = p.c.Call(ctx, "UserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserFollow(ctx context.Context, req *user.UserFollowRequest) (r *user.UserFollowResponse, err error) {
	var _args user.UserServiceUserFollowArgs
	_args.Req = req
	var _result user.UserServiceUserFollowResult
	if err = p.c.Call(ctx, "UserFollow", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserUnfollow(ctx context.Context, req *user.UserUnfollowRequest) (r *user.UserUnfollowResponse, err error) {
	var _args user.UserServiceUserUnfollowArgs
	_args.Req = req
	var _result user.UserServiceUserUnfollowResult
	if err = p.c.Call(ctx, "UserUnfollow", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFollowList(ctx context.Context, req *user.FollowListRequest) (r *user.FollowListResponse, err error) {
	var _args user.UserServiceGetFollowListArgs
	_args.Req = req
	var _result user.UserServiceGetFollowListResult
	if err = p.c.Call(ctx, "GetFollowList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFollowerList(ctx context.Context, req *user.FollowerListRequest) (r *user.FollowerListResponse, err error) {
	var _args user.UserServiceGetFollowerListArgs
	_args.Req = req
	var _result user.UserServiceGetFollowerListResult
	if err = p.c.Call(ctx, "GetFollowerList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFriendList(ctx context.Context, req *user.FriendListRequest) (r *user.FriendListResponse, err error) {
	var _args user.UserServiceGetFriendListArgs
	_args.Req = req
	var _result user.UserServiceGetFriendListResult
	if err = p.c.Call(ctx, "GetFriendList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) IsFriend(ctx context.Context, req *user.IsFriendRequest) (r *user.IsFriendResponse, err error) {
	var _args user.UserServiceIsFriendArgs
	_args.Req = req
	var _result user.UserServiceIsFriendResult
	if err = p.c.Call(ctx, "IsFriend", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CompGetUser(ctx context.Context, req *user.CompGetUserRequest) (r *user.CompGetUserResponse, err error) {
	var _args user.UserServiceCompGetUserArgs
	_args.Req = req
	var _result user.UserServiceCompGetUserResult
	if err = p.c.Call(ctx, "CompGetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CompMGetUser(ctx context.Context, req *user.CompMGetUserRequest) (r *user.CompMGetUserResponse, err error) {
	var _args user.UserServiceCompMGetUserArgs
	_args.Req = req
	var _result user.UserServiceCompMGetUserResult
	if err = p.c.Call(ctx, "CompMGetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RelationAction(ctx context.Context, req *user.RelationActionRequest) (r *user.RelationActionResponse, err error) {
	var _args user.UserServiceRelationActionArgs
	_args.Req = req
	var _result user.UserServiceRelationActionResult
	if err = p.c.Call(ctx, "RelationAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
