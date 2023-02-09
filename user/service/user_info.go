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

func UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	resp = user.NewUserInfoResponse()
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserInfo")
	defer span.Finish()

	var q = query.Use(dal.DB.Debug())
	userDao := q.User.WithContext(ctx)
	gormUser, err := userDao.FindByUserID(uint(req.UserId))

	// 如果查询失败
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.BaseResp = util.PackBaseResp(errno.UserNotExist)
		} else {
			resp.BaseResp = util.PackBaseResp(err)
		}

		return resp, errno.UserNotExist
	}

	followCount := int64(gormUser.FollowCount)
	followerCount := int64(gormUser.FollowerCount)

	resp.User = &user.User{
		Id:            int64(gormUser.ID),
		Name:          gormUser.Name,
		FollowCount:   &followCount,
		FollowerCount: &followerCount,
	}
	resp.BaseResp = util.PackBaseResp(errno.Success)
	return resp, nil
}
