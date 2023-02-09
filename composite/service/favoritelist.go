package service

import (
	"context"

	"github.com/41197-yhkt/tiktok/composite/gen/dal/model"
	"github.com/41197-yhkt/tiktok/composite/kitex_gen/composite"
	"github.com/41197-yhkt/tiktok/composite/pack"
	"github.com/41197-yhkt/tiktok/composite/rpc"
	"github.com/41197-yhkt/tiktok/video/kitex_gen/video"

	"github.com/41197-yhkt/tiktok/pkg/errno"
	"github.com/41197-yhkt/tiktok/user/kitex_gen/user"
)

type FavoriteListService struct {
	ctx context.Context
}

func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{ctx: ctx}
}

func (s *FavoriteListService) FavoriteList(req *composite.BasicFavoriteListRequest) ([]*composite.Video, error) {
	userFavoriteDatabase := q.UserFavorite.WithContext(s.ctx)

	// 1. 从 user_favorites 中根据 query_id 查 video_id
	var userFavorites []*model.UserFavorite
	userFavorites, err := userFavoriteDatabase.FindByUserid(req.QueryId)
	if err != nil {
		return nil, errno.UserNotExist
	}

	// 2. 对于每个 video_id
	// TODO: 接到 user 和 video 服务上
	videoIds, authorIds := pack.VideoAndVideoAuthorIds(userFavorites)
	authors, err := rpc.MGetUser(s.ctx, &user.CompMGetUserRequest{
		UserId:        req.UserId,
		TargetUsersId: authorIds,
	})
	videos, err := rpc.MGetVideo(s.ctx, &video.MGetVideoRequest{
		UserId:         req.UserId,
		TargetVideosId: videoIds,
	})
	if err != nil {
		return nil, err
	}

	res := pack.Videos(videos, authors)
	return res, nil
}
