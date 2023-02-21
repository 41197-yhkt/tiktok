package service

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/41197-yhkt/tiktok/user/kitex_gen/user"
	video "github.com/41197-yhkt/tiktok/video/kitex_gen/video"
	"github.com/41197-yhkt/tiktok/video/rpc"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// service/publishlist.go
type PublishListService struct {
	ctx context.Context
}

func NewPublishListService(ctx context.Context) *PublishListService {
	return &PublishListService{ctx: ctx}
}

func (s *PublishListService) PublishList(req *video.DouyinPublishListRequest) ([]*video.Video, error) {
	hlog.Info("PublishList: req=", req)
	//videoDatabase := q.Video.WithContext(s.ctx)
	//client为阿里云oss对象
	client, err := oss.New("oss-cn-hangzhou.aliyuncs.com", "LTAI5tMJ1hvEXziXradJWqmt", "vJ1MFiiqHbmZQSOQkLbD0EDLhjgagD")
	if err != nil {
		log.Panic(err)
	}
	//选择视频bucket
	bucket, err := client.Bucket("41197-tiktok-bucket")
	if err != nil {
		log.Panic(err)
	}
	videoDatabase := q.Video.WithContext(s.ctx)
	userFavoriteDatabase := q.UserFavorite.WithContext(s.ctx)
	//userDatabase := q.User.WithContext(s.ctx)

	// 先根据 user_id 选出 videos
	var isFavorite bool
	videos, err := videoDatabase.FindByAuthorId(int(req.TargetId))
	fmt.Println(videos)

	// 根据 video_id 查 Video
	res := []*video.Video{}
	for _, vd := range videos {
		author, err := rpc.GetUser(s.ctx, &user.CompGetUserRequest{UserId: req.UserId, TargetUserId: vd.AuthorId})
		if err != nil {
			return nil, err
		}

		// 查询点赞数目
		var favoriteCount int64
		favoriteCount, err = userFavoriteDatabase.CountByVideoid(int64(vd.ID))

		// TODO: 查询评论数
		var commentCount int64
		commentCount, err = strconv.ParseInt(vd.CommentCount, 10, 64)

		// 查询自己是不是也点了赞
		err = userFavoriteDatabase.FindByUseridAndVideoid(req.UserId, int64(vd.ID))
		if err != nil {
			isFavorite = false
		} else {
			isFavorite = true
		}

		playurl, err := bucket.SignURL(vd.PlayUrl, oss.HTTPGet, 30)
		if err != nil {
			panic(err)
		}

		coverurl, err := bucket.SignURL("cover.png", oss.HTTPGet, 30)
		if err != nil {
			panic(err)
		}

		// 封装
		res = append(res, &video.Video{
			Id: int64(vd.ID),
			Author: &video.User{
				Id:            author.Id,
				Name:          author.Name,
				FollowCount:   author.FollowCount,
				FollowerCount: author.FollowerCount,
				IsFollow:      author.IsFollow,
			},
			PlayUrl:       playurl,
			CoverUrl:      coverurl,
			FavoriteCount: favoriteCount,
			CommentCount:  commentCount,
			IsFavorite:    isFavorite,
			Title:         vd.Title,
		})
	}

	return res, nil
}
