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

func GetFriendList(ctx context.Context, req *user.FriendListRequest) (resp *user.FriendListResponse, err error) {
	resp = user.NewFriendListResponse()
	span, ctx := opentracing.StartSpanFromContext(ctx, "FriendList")
	defer span.Finish()

	var q = query.Use(dal.DB.Debug())
	userRelationDao := q.UserRelation.WithContext(ctx)
	userDao := q.User.WithContext(ctx)
	userId := req.UserId

	//找到user的关注列表，得到的userFollow_list的follow_from均为userId
	userFollow_list, err := userRelationDao.FindByFollowFrom(uint(userId))

	if err != nil {
		resp.BaseResp = util.PackBaseResp(err)
		return
	}

	userList := make([]*user.User, 0)
	for _, v := range userFollow_list {
		//判断user的关注列表有没有跟user是朋友的
		boolResp, erro := IsFriend(ctx, userId, int64(v.FollowTo))
		if erro != nil {
			resp.BaseResp = util.PackBaseResp(err)
			return
		}
		//如果是朋友，则直接找出该follow_user信息，并加入到resp user_list中
		if boolResp.IsFriend {
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
	}

	resp.UserList = userList
	resp.BaseResp = util.PackBaseResp(nil)
	return resp, nil
}
