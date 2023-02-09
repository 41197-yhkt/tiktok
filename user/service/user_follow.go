package service

import (
	"context"
	"errors"
	"github.com/41197-yhkt/pkg/errno"
	"github.com/41197-yhkt/tiktok-user/dao/dal"
	"github.com/41197-yhkt/tiktok-user/dao/dal/model"
	"github.com/41197-yhkt/tiktok-user/dao/dal/query"
	"github.com/41197-yhkt/tiktok-user/kitex_gen/user"
	"github.com/41197-yhkt/tiktok-user/util"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

// UserFollow 用户关注，就算关注多次，数据库实际也只保存一条记录昂
func UserFollow(ctx context.Context, req *user.UserFollowRequest) (resp *user.UserFollowResponse, err error) {
	resp = user.NewUserFollowResponse()
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserInfo")
	defer span.Finish()

	var q = query.Use(dal.DB.Debug())
	userRelationDao := q.UserRelation.WithContext(ctx)

	followFrom := uint(req.GetFollowFrom())
	followTo := uint(req.GetFollowTo())

	_, err = userRelationDao.FindByFollowFromAndFollowTo(followFrom, followTo)
	// 如果当前查询的关注存在
	if err == nil {
		resp.BaseResp = util.PackBaseResp(errno.UserFollowRelationExistErr)
		return resp, errno.UserFollowRelationExistErr
	}

	// 当前的关注关系不存在
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		newUserRelation := model.UserRelation{
			FollowFrom: followFrom,
			FollowTo:   followTo,
		}

		err = userRelationDao.Create(&newUserRelation)

		// 创建关注关系失败
		if err != nil {
			resp.BaseResp = util.PackBaseResp(err)
			return resp, err
		}
		// 正常返回
		resp.BaseResp = util.PackBaseResp(errno.Success)
		return resp, nil
	} else {
		// 上述查询出现其他错误
		resp.BaseResp = util.PackBaseResp(err)
		return resp, err
	}
}

// UserUnfollow 用户取消关注，由于存在软删除
func UserUnfollow(ctx context.Context, req *user.UserUnfollowRequest) (resp *user.UserUnfollowResponse, err error) {
	resp = user.NewUserUnfollowResponse()
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserInfo")
	defer span.Finish()

	var q = query.Use(dal.DB.Debug())
	userRelationDao := q.UserRelation.WithContext(ctx)

	// 先把对应的FollowRelation查询出来
	followFrom := uint(req.GetFollowFrom())
	followTo := uint(req.GetFollowTo())
	userRelation, sErr := userRelationDao.FindByFollowFromAndFollowTo(followFrom, followTo)

	if sErr != nil {
		// 要删除的记录不存在
		if errors.Is(sErr, gorm.ErrRecordNotFound) {
			resp.BaseResp = util.PackBaseResp(errno.UserFollowRelationNotExistErr)
			return resp, errno.UserFollowRelationNotExistErr
		} else {
			resp.BaseResp = util.PackBaseResp(sErr)
			return resp, sErr
		}
	}

	_, err = userRelationDao.Delete(&userRelation)
	if err != nil {
		resp.BaseResp = util.PackBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = util.PackBaseResp(errno.Success)
	return resp, nil
}
