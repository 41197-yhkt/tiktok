// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	user "github.com/41197-yhkt/tiktok/user/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	UserRegister(ctx context.Context, req *user.UserRegisterRequest, callOptions ...callopt.Option) (r *user.UserRegisterResponse, err error)
	UserLogin(ctx context.Context, req *user.UserLoginRequest, callOptions ...callopt.Option) (r *user.UserLoginResponse, err error)
	UserInfo(ctx context.Context, req *user.UserInfoRequest, callOptions ...callopt.Option) (r *user.UserInfoResponse, err error)
	UserFollow(ctx context.Context, req *user.UserFollowRequest, callOptions ...callopt.Option) (r *user.UserFollowResponse, err error)
	UserUnfollow(ctx context.Context, req *user.UserUnfollowRequest, callOptions ...callopt.Option) (r *user.UserUnfollowResponse, err error)
	GetFollowList(ctx context.Context, req *user.FollowListRequest, callOptions ...callopt.Option) (r *user.FollowListResponse, err error)
	GetFollowerList(ctx context.Context, req *user.FollowerListRequest, callOptions ...callopt.Option) (r *user.FollowerListResponse, err error)
	GetFriendList(ctx context.Context, req *user.FriendListRequest, callOptions ...callopt.Option) (r *user.FriendListResponse, err error)
	IsFriend(ctx context.Context, req *user.IsFriendRequest, callOptions ...callopt.Option) (r *user.IsFriendResponse, err error)
	CompGetUser(ctx context.Context, req *user.CompGetUserRequest, callOptions ...callopt.Option) (r *user.CompGetUserResponse, err error)
	CompMGetUser(ctx context.Context, req *user.CompMGetUserRequest, callOptions ...callopt.Option) (r *user.CompMGetUserResponse, err error)
	RelationAction(ctx context.Context, req *user.RelationActionRequest, callOptions ...callopt.Option) (r *user.RelationActionResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) UserRegister(ctx context.Context, req *user.UserRegisterRequest, callOptions ...callopt.Option) (r *user.UserRegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserRegister(ctx, req)
}

func (p *kUserServiceClient) UserLogin(ctx context.Context, req *user.UserLoginRequest, callOptions ...callopt.Option) (r *user.UserLoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserLogin(ctx, req)
}

func (p *kUserServiceClient) UserInfo(ctx context.Context, req *user.UserInfoRequest, callOptions ...callopt.Option) (r *user.UserInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserInfo(ctx, req)
}

func (p *kUserServiceClient) UserFollow(ctx context.Context, req *user.UserFollowRequest, callOptions ...callopt.Option) (r *user.UserFollowResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserFollow(ctx, req)
}

func (p *kUserServiceClient) UserUnfollow(ctx context.Context, req *user.UserUnfollowRequest, callOptions ...callopt.Option) (r *user.UserUnfollowResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserUnfollow(ctx, req)
}

func (p *kUserServiceClient) GetFollowList(ctx context.Context, req *user.FollowListRequest, callOptions ...callopt.Option) (r *user.FollowListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowList(ctx, req)
}

func (p *kUserServiceClient) GetFollowerList(ctx context.Context, req *user.FollowerListRequest, callOptions ...callopt.Option) (r *user.FollowerListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowerList(ctx, req)
}

func (p *kUserServiceClient) GetFriendList(ctx context.Context, req *user.FriendListRequest, callOptions ...callopt.Option) (r *user.FriendListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFriendList(ctx, req)
}

func (p *kUserServiceClient) IsFriend(ctx context.Context, req *user.IsFriendRequest, callOptions ...callopt.Option) (r *user.IsFriendResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IsFriend(ctx, req)
}

func (p *kUserServiceClient) CompGetUser(ctx context.Context, req *user.CompGetUserRequest, callOptions ...callopt.Option) (r *user.CompGetUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CompGetUser(ctx, req)
}

func (p *kUserServiceClient) CompMGetUser(ctx context.Context, req *user.CompMGetUserRequest, callOptions ...callopt.Option) (r *user.CompMGetUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CompMGetUser(ctx, req)
}

func (p *kUserServiceClient) RelationAction(ctx context.Context, req *user.RelationActionRequest, callOptions ...callopt.Option) (r *user.RelationActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationAction(ctx, req)
}
