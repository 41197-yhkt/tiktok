package service

import (
	"context"
	"time"

	"github.com/41197-yhkt/tiktok/tiktok-composite/kitex_gen/composite"
	"github.com/41197-yhkt/tiktok/tiktok-composite/pack"
	"github.com/41197-yhkt/tiktok/tiktok-composite/rpc"
	"github.com/41197-yhkt/tiktok/tiktok-user/kitex_gen/user"
	"github.com/41197-yhkt/tiktok/tiktok-video/kitex_gen/video"
)

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{ctx: ctx}
}

func (s *FeedService) Feed(req *composite.BasicFeedRequest) ([]*composite.Video, int64, error) {
	videoDatabase := q.Video.WithContext(s.ctx)

	// 1. 查看 req.Last_time 是否存在，若存在使用 req.Last_time，否则使用 time.Now()
	lastest_time := time.Now()
	if req.LastestTime != nil {
		lastest_time = time.Unix(*req.LastestTime, 0)
	}

	// 2. 查找 video 表中更新时间小于 last_time 的最大 limit 个视频
	eligibleVideos, err := videoDatabase.FindByUpdatedtime(lastest_time, 30)
	if err != nil {
		return nil, lastest_time.Unix(), err
	}
	// 3.1. video 表中查视频信息
	// TODO: 接入 video 服务
	videoIds := pack.VideoIds(eligibleVideos)
	videos, err := rpc.MGetVideo(s.ctx, &video.MGetVideoRequest{
		UserId:         req.UserId,
		TargetVideosId: videoIds,
	})

	// 3.2. user 表中查作者信息
	authorIds := pack.VideoAuthorIds(eligibleVideos)
	authors, err := rpc.MGetUser(s.ctx, &user.CompMGetUserRequest{
		UserId:        req.UserId,
		TargetUsersId: authorIds,
	})
	if err != nil {
		return nil, lastest_time.Unix(), err
	}

	res := pack.Videos(videos, authors)
	if len(eligibleVideos) > 0 {
		return res, eligibleVideos[0].UpdatedAt.Unix(), nil
	}
	return res, lastest_time.Unix(), nil
}
