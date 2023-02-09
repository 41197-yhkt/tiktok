package service

import (
	"context"
	"errors"
	"github.com/41197-yhkt/pkg/errno"
	"github.com/41197-yhkt/tiktok-user/dao/dal"
	"github.com/41197-yhkt/tiktok-user/dao/dal/query"
	"github.com/41197-yhkt/tiktok-user/kitex_gen/user"
	"github.com/41197-yhkt/tiktok-user/util"
	"gorm.io/gorm"
)

func CompMGetUser(ctx context.Context, req *user.CompMGetUserRequest) (resp *user.CompMGetUserResponse, err error) {
	resp = user.NewCompMGetUserResponse()
	var q = query.Use(dal.DB.Debug())
	userRelationDao := q.UserRelation.WithContext(ctx)
	userDao := q.User.WithContext(ctx)

	// construct targetUserIDs
	followFrom := uint(req.UserId)
	targetUserIDs := make([]uint, len(req.TargetUsersId))
	for i, v := range req.TargetUsersId {
		targetUserIDs[i] = uint(v)
	}

	// 获取当前用户的关注列表
	followList, err := userRelationDao.FindByFollowFrom(followFrom)
	if err != nil {
		resp.BaseResp = util.PackBaseResp(err)
		return
	}

	userList := make([]*user.User, 0)
	for _, targetUserID := range targetUserIDs {
		targetUser, err := userDao.FindByUserID(targetUserID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				resp.BaseResp = util.PackBaseResp(errno.UserNotExist)
			} else {
				resp.BaseResp = util.PackBaseResp(err)
			}
			return resp, err
		}
		followCount := int64(targetUser.FollowCount)
		followerCount := int64(targetUser.FollowerCount)

		packedUser := &user.User{
			Id:            int64(targetUser.ID),
			Name:          targetUser.Name,
			FollowCount:   &followCount,
			FollowerCount: &followerCount,
		}

		// 如果当前目标用户处于关注列表中
		for _, followRelation := range followList {
			if followRelation.FollowTo == targetUserID {
				packedUser.IsFollow = true
			}
		}

		userList = append(userList, packedUser)
	}

	resp.UserList = userList
	resp.BaseResp = util.PackBaseResp(errno.Success)
	return
}
