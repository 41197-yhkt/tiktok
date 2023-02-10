package service

import (
	"context"
	"log"
	"strconv"

	video "github.com/41197-yhkt/tiktok/video/kitex_gen/video"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// service/GetVideo.go
type GetVideoService struct {
	ctx context.Context
}

func NewGetVideoService(ctx context.Context) *GetVideoService {
	return &GetVideoService{ctx: ctx}
}

func (s *GetVideoService) GetVideo(req *video.GetVideoRequest) (*video.Video, error) {
	//videoDatabase := q.Video.WithContext(s.ctx)
	//client为阿里云oss对象
	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "LTAI5tGdrFczu9cP7RX8LgrC", "I0P6eEUAk740O5jM1VLbvfePs5yGAf")
	if err != nil {
		log.Panic(err)
	}
	//选择视频bucket
	bucket, err := client.Bucket("video-bucket0")
	if err != nil {
		log.Panic(err)
	}

	userFavoriteDatabase := q.UserFavorite.WithContext(s.ctx)
	videoDatabase := q.Video.WithContext(s.ctx)

	// 先根据 user_id 选出 videos
	var isFavorite bool
	var res *video.Video
	videos, err := videoDatabase.FindByID(req.TargetVideoId)
	if err != nil {
		return nil, err
	}

	// 查询点赞数目
	favoriteCount, err := userFavoriteDatabase.CountByVideoid(int64(videos.ID))
	if err != nil {
		return nil, err
	}

	// TODO: 查询评论数
	var commentCount int64
	commentCount, err = strconv.ParseInt(videos.CommentCount, 10, 64)

	// 查询自己是不是也点了赞
	err = userFavoriteDatabase.FindByUseridAndVideoid(req.UserId, req.TargetVideoId)
	if err != nil {
		isFavorite = false
	} else {
		isFavorite = true
	}

	playurl, err := bucket.SignURL(videos.Title, oss.HTTPGet, 30)
	if err != nil {
		panic(err)
	}

	coverurl, err := bucket.SignURL("cover.png", oss.HTTPGet, 30)
	if err != nil {
		panic(err)
	}

	// 封装
	res = &video.Video{
		Id: req.TargetVideoId,
		Author: &video.User{
			Id: videos.AuthorId,
		},
		PlayUrl:       playurl,
		CoverUrl:      coverurl,
		FavoriteCount: favoriteCount,
		CommentCount:  commentCount,
		IsFavorite:    isFavorite,
		Title:         videos.Title,
	}
	// fmt.Println(res)
	return res, nil
}
