package service

import (
	"context"
	"errors"

	"github.com/41197-yhkt/pkg/errno"
	"github.com/41197-yhkt/tiktok-user/dao/dal"
	"github.com/41197-yhkt/tiktok-user/dao/dal/query"
	"github.com/41197-yhkt/tiktok-user/kitex_gen/user"
	"github.com/41197-yhkt/tiktok-user/util"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

func GetFollowList(ctx context.Context, req *user.FollowListRequest) (resp *user.FollowListResponse, err error) {
	resp = user.NewFollowListResponse()
	span, ctx := opentracing.StartSpanFromContext(ctx, "FollowList")
	defer span.Finish()
	var q = query.Use(dal.DB.Debug())
	userRelationDao := q.UserRelation.WithContext(ctx)
	userDao := q.User.WithContext(ctx)
	userID := req.UserId
	//找到user的关注列表
	userRelationList, err := userRelationDao.FindByFollowFrom(uint(userID))
	if err != nil {
		resp.BaseResp = util.PackBaseResp(err)
		return
	}
	//根据user的关注列表ID找到对应UserInfo
	userList := make([]*user.User, 0)
	for _, v := range userRelationList {
		followUser, sErr := userDao.FindByUserID(v.FollowTo)
		if sErr != nil {
			if errors.Is(sErr, gorm.ErrRecordNotFound) {
				resp.BaseResp = util.PackBaseResp(errno.UserNotExist)
			} else {
				resp.BaseResp = util.PackBaseResp(sErr)
			}
			return resp, sErr
		}
		followCount := int64(followUser.FollowCount)
		followerCount := int64(followUser.FollowerCount)
		userList = append(userList, &user.User{
			Id:            int64(followUser.ID),
			Name:          followUser.Name,
			FollowCount:   &followCount,
			FollowerCount: &followerCount,
			IsFollow:      false,
		})
	}

	resp.UserList = userList
	resp.BaseResp = util.PackBaseResp(errno.Success)
	return resp, nil

}
