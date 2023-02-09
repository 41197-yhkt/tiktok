package service

import (
	"context"

	"github.com/41197-yhkt/tiktok-user/dao/dal"
	"github.com/41197-yhkt/tiktok-user/dao/dal/query"
	"github.com/41197-yhkt/tiktok-user/kitex_gen/user"
	"github.com/41197-yhkt/tiktok-user/util"
	"github.com/opentracing/opentracing-go"
)

func IsFriend(ctx context.Context, userId, toUserId int64) (resp *user.IsFriendResponse, err error) {
	resp = user.NewIsFriendResponse()
	span, ctx := opentracing.StartSpanFromContext(ctx, "IsFriend")
	defer span.Finish()
	var q = query.Use(dal.DB.Debug())
	UserRelationDao := q.UserRelation.WithContext(ctx)

	//找到user1的关注列表
	userIdFollowList, err := UserRelationDao.FindByFollowFrom(uint(userId))
	if err != nil {
		resp.BaseResp = util.PackBaseResp(err)
		return
	}
	//找到user2的关注列表
	toUserIdFollowList, err := UserRelationDao.FindByFollowFrom(uint(toUserId))
	if err != nil {
		resp.BaseResp = util.PackBaseResp(err)
		return
	}
	var flag1, flag2 bool
	//判断user1的关注列表是否有user2
	for _, v := range userIdFollowList {
		if toUserId == int64(v.FollowTo) {
			flag1 = true
			break
		}
	}
	//判断user2的关注列表是否有user1
	for _, v := range toUserIdFollowList {
		if userId == int64(v.FollowTo) {
			flag2 = true
			break
		}
	}
	resp.IsFriend = flag1 && flag2
	resp.BaseResp = util.PackBaseResp(nil)
	return resp, nil
}
