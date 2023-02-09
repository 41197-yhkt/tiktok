package main

import (
	"context"

	user "github.com/41197-yhkt/tiktok-user/kitex_gen/user"
	"github.com/41197-yhkt/tiktok-user/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	return service.UserRegister(ctx, req)
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	return service.UserLogin(ctx, req)
}

// UserFollow implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserFollow(ctx context.Context, req *user.UserFollowRequest) (resp *user.UserFollowResponse, err error) {
	return service.UserFollow(ctx, req)
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	return service.UserInfo(ctx, req)
}

// UserUnfollow implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserUnfollow(ctx context.Context, req *user.UserUnfollowRequest) (resp *user.UserUnfollowResponse, err error) {
	return service.UserUnfollow(ctx, req)
}

// GetFollowList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFollowList(ctx context.Context, req *user.FollowListRequest) (resp *user.FollowListResponse, err error) {
	return service.GetFollowList(ctx, req)
}

// GetFollowerList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFollowerList(ctx context.Context, req *user.FollowerListRequest) (resp *user.FollowerListResponse, err error) {
	return service.GetFollowerList(ctx, req)
}

// GetFriendList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFriendList(ctx context.Context, req *user.FriendListRequest) (resp *user.FriendListResponse, err error) {
	return service.GetFriendList(ctx, req)
}

// IsFriend implements the UserServiceImpl interface.
func (s *UserServiceImpl) IsFriend(ctx context.Context, req *user.IsFriendRequest) (resp *user.IsFriendResponse, err error) {
	return service.IsFriend(ctx, req.UserId, req.ToUserId)
}

// CompGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CompGetUser(ctx context.Context, req *user.CompGetUserRequest) (resp *user.CompGetUserResponse, err error) {
	return service.CompGetUser(ctx, req)
}

// CompMGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CompMGetUser(ctx context.Context, req *user.CompMGetUserRequest) (resp *user.CompMGetUserResponse, err error) {
	return service.CompMGetUser(ctx, req)
}
