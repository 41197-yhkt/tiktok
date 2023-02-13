package service

import (
	"bytes"
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/41197-yhkt/tiktok/video/gen/dal/model"
	video "github.com/41197-yhkt/tiktok/video/kitex_gen/video"
	"github.com/cloudwego/hertz/pkg/common/hlog"

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
	hlog.Info("PublishAction: fileName=", req.Filename, "title: ", req.Title)

	//client为阿里云oss对象
	client, err := oss.New("oss-cn-hangzhou.aliyuncs.com", "LTAI5tMJ1hvEXziXradJWqmt", "vJ1MFiiqHbmZQSOQkLbD0EDLhjgagD")
	if err != nil {
		hlog.Info("PublishAction: err=", err.Error())
		return err
	}
	//选择视频bucket
	bucket, err := client.Bucket("41197-tiktok-bucket")
	if err != nil {
		hlog.Info("PublishAction: err=", err.Error())
		return err
	}

	title := strconv.Itoa(int(time.Now().Unix())) + req.Filename
	//上传视频流
	err = bucket.PutObject(title, bytes.NewReader(req.Data))
	if err != nil {
		hlog.Info("PublishAction: err=", err.Error())
		return
	}

	videoDatabase := q.Video.WithContext(s.ctx)

	//获取视频流URL
	// playurl, err := bucket.SignURL(title, oss.HTTPGet, 30)
	// if err != nil {
	// 	hlog.Info("PublishAction: err=", err.Error())
	// 	return
	// }

	//获取固定封面URL
	playurl_cover, err := bucket.SignURL("cover.png", oss.HTTPGet, 30)
	if err != nil {
		hlog.Info("PublishAction: err=", err.Error())
		return
	}

	//构建数据库结构体
	videodata := &model.Video{
		AuthorId:      req.Author,
		PlayUrl:       title,
		CoverUrl:      playurl_cover,
		FavoriteCount: "0",
		CommentCount:  "0",
		Title:         req.Title,
	}

	err = videoDatabase.WithContext(s.ctx).Create(videodata)
	if err != nil {
		hlog.Info("PublishAction: err=", err.Error())
		return
	}
	return nil
}
