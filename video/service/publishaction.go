package service

import (
	"bytes"
	"context"
	"sync"
	"time"

	"github.com/41197-yhkt/tiktok/video/gen/dal/model"
	video "github.com/41197-yhkt/tiktok/video/kitex_gen/video"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var once sync.Once
var client *oss.Client
var bucket *oss.Bucket

func init() {
	once.Do(func() {

	})
}

type PublishActionService struct {
	ctx context.Context
}

func NewPublishActionService(ctx context.Context) *PublishActionService {
	return &PublishActionService{ctx: ctx}
}

func (s *PublishActionService) PublishAction(req *video.DouyinPublishActionRequest) (err error) {
	//client为阿里云oss对象
	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "LTAI5tGdrFczu9cP7RX8LgrC", "I0P6eEUAk740O5jM1VLbvfePs5yGAf")
	if err != nil {
		return err
	}
	//选择视频bucket
	bucket, err := client.Bucket("video-bucket0")
	if err != nil {
		return err
	}

	title := req.Title + string(rune(time.Now().Unix()))

	//上传视频流
	err = bucket.PutObject(title, bytes.NewReader(req.Data))
	if err != nil {
		return
	}

	videoDatabase := q.Video.WithContext(s.ctx)

	//获取视频流URL
	playurl, err := bucket.SignURL(title, oss.HTTPGet, 30)
	if err != nil {
		return
	}

	//获取固定封面URL
	playurl_cover, err := bucket.SignURL("cover.png", oss.HTTPGet, 30)
	if err != nil {
		return
	}

	//构建数据库结构体
	videodata := &model.Video{
		AuthorId:      req.Author,
		PlayUrl:       playurl,
		CoverUrl:      playurl_cover,
		FavoriteCount: "0",
		CommentCount:  "0",
		Title:         title,
	}

	err = videoDatabase.WithContext(s.ctx).Create(videodata)
	if err != nil {
		return
	}
	return nil
}
