package service

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/41197-yhkt/tiktok-user/kitex_gen/user"
	dal "github.com/41197-yhkt/tiktok-video/gen/dal"
	"github.com/41197-yhkt/tiktok-video/gen/dal/model"
	video "github.com/41197-yhkt/tiktok-video/kitex_gen/video"
	"github.com/41197-yhkt/tiktok-video/rpc"

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
	//userDatabase := q.User.WithContext(s.ctx)

	// 先根据 user_id 选出 videos
	var videos []*model.Video
	var isFavorite bool

	dal.DB.WithContext(s.ctx).Where("author_id = ?", req.UserId).Find(&videos)
	fmt.Println(videos)

	// 根据 video_id 查 Video
	res := []*video.Video{}
	for _, vd := range videos {
		// video, err := videoDatabase.FindByID(int64(vd.Id))
		// if err != nil {
		//     panic(err)
		// }
		author, err := rpc.GetUser(s.ctx, &user.CompGetUserRequest{UserId: req.UserId, TargetUserId: vd.AuthorId})
		if err != nil {
			return nil, err
		}

		// 查询点赞数目
		var favoriteCount int64
		dal.DB.WithContext(s.ctx).Where("author_id = ?", vd.Id).Count(&favoriteCount)

		// TODO: 查询评论数
		var commentCount int64
		commentCount, err = strconv.ParseInt(vd.CommentCount, 10, 64)

		// 查询自己是不是也点了赞

		err = userFavoriteDatabase.WithContext(s.ctx).FindByUseridAndVideoid(req.UserId, int64(vd.Id))
		if err != nil {
			isFavorite = false
		} else {
			isFavorite = true
		}

		playurl, err := bucket.SignURL(vd.Title, oss.HTTPGet, 30)
		if err != nil {
			panic(err)
		}

		coverurl, err := bucket.SignURL("cover.png", oss.HTTPGet, 30)
		if err != nil {
			panic(err)
		}

		// 封装
		res = append(res, &video.Video{
			Id: int64(vd.Id),
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
