package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/41197-yhkt/tiktok-video/kitex_gen/video"
	"github.com/41197-yhkt/tiktok-video/kitex_gen/video/douyinservice"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	cli "github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var VideoClient douyinservice.Client

func setupClient() douyinservice.Client {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}

	c, err := douyinservice.NewClient(
		"video",
		cli.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	return c
}

func TestOss(t *testing.T) {
	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "LTAI5tGdrFczu9cP7RX8LgrC", "I0P6eEUAk740O5jM1VLbvfePs5yGAf")
	if err != nil {
		t.Error(err)
	}
	//选择视频bucket
	bucket, err := client.Bucket("video-bucket0")
	if err != nil {
		t.Error(err)
	}
	playurl, err := bucket.SignURL("helloworld", oss.HTTPGet, 30)
	if err != nil {
		panic(err)
	}
	fmt.Println(playurl)
}

func TestGetVedio(t *testing.T) {
	videoClient := setupClient()
	ctx := context.Background()

	var req *video.GetVideoRequest
	req = &video.GetVideoRequest{
		UserId:        1,
		TargetVideoId: 1,
	}
	video, err := videoClient.DouyinGetVideoMethod(ctx, req)
	if err != nil {
		t.Error(err)
	}
	t.Log(video.Video)
}
