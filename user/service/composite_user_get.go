package service

import (
	"context"
	"github.com/41197-yhkt/pkg/errno"
	"github.com/41197-yhkt/tiktok-user/dao/dal"
	"github.com/41197-yhkt/tiktok-user/dao/dal/query"
	"github.com/41197-yhkt/tiktok-user/kitex_gen/user"
	"github.com/41197-yhkt/tiktok-user/util"
)

func CompGetUser(ctx context.Context, req *user.CompGetUserRequest) (resp *user.CompGetUserResponse, err error) {
	resp = user.NewCompGetUserResponse()

	// get target_user_info
	targetUserID := req.GetTargetUserId()
	getUserReq := user.NewUserInfoRequest()
	getUserReq.UserId = targetUserID
	targetUser, err := UserInfo(ctx, getUserReq)
	if err != nil {
		resp.BaseResp = util.PackBaseResp(err)
		return
	}

	// judge follow relation
	var q = query.Use(dal.DB.Debug())
	followFrom := uint(req.UserId)
	targetFollowTo := uint(req.TargetUserId)
	userRelationDao := q.UserRelation.WithContext(ctx)
	userRelationList, err := userRelationDao.FindByFollowFrom(followFrom)
	if err != nil {
		resp.BaseResp = util.PackBaseResp(err)
		return
	}

	for _, userRelation := range userRelationList {
		if userRelation.FollowTo == targetFollowTo {
			targetUser.User.IsFollow = true
		}
	}

	resp.User = targetUser.User
	resp.BaseResp = util.PackBaseResp(errno.Success)
	return
}
