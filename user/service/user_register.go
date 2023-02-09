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

func UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	resp = user.NewUserRegisterResponse()
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRegister")
	defer span.Finish()

	var q = query.Use(dal.DB)
	userDao := q.User.WithContext(ctx)
	_, err = userDao.FindByUserName(req.Username)

	if err == nil {
		// username对应的用户已经存在，请重新注册
		resp.BaseResp = util.PackBaseResp(errno.UserAlreadyExistErr)
		return resp, errno.UserAlreadyExistErr
	}

	// 当前注册的用户不存在
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		// 对于用户密码进行加密处理
		rawPassword := req.Password
		password, sErr := util.EncryptPasswd(rawPassword)
		if sErr != nil {
			resp.BaseResp = util.PackBaseResp(sErr)
			return resp, sErr
		}

		newUser := &model.User{
			Name:     req.Username,
			Password: password,
		}

		// 创建新用户失败
		createRes := dal.DB.Create(newUser)
		if createRes.Error != nil {
			resp.BaseResp = util.PackBaseResp(createRes.Error)
			return resp, createRes.Error
		}

		// 正常返回
		resp.UserId = int64(newUser.ID)
		resp.BaseResp = util.PackBaseResp(errno.Success)
		return resp, nil
	} else {
		resp.BaseResp = util.PackBaseResp(err)
		return resp, err
	}
}
