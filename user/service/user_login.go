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

func UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	resp = user.NewUserLoginResponse()
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserLogin")
	defer span.Finish()

	var q = query.Use(dal.DB.Debug())
	userDao := q.User.WithContext(ctx)
	gormUser, sErr := userDao.FindByUserName(req.Username)

	// 如果记录不存在
	if errors.Is(sErr, gorm.ErrRecordNotFound) {
		resp.BaseResp = util.PackBaseResp(errno.UserNotExist)
		return resp, errno.UserNotExist
	}

	// 如果记录存在
	pwdCmpPass, sErr := util.ComparePasswd(gormUser.Password, req.Password)
	// 密码比对出现问题
	if sErr != nil {
		resp.BaseResp = util.PackBaseResp(sErr)
		return resp, sErr
	}
	// 密码不匹配
	if !pwdCmpPass {
		resp.BaseResp = util.PackBaseResp(errno.UserPwdErr)
		return resp, errno.UserPwdErr
	}

	resp.BaseResp = util.PackBaseResp(errno.Success)
	resp.UserId = int64(gormUser.ID)
	return resp, nil
}
