package rpc

import (
	"context"
	"time"

	"github.com/41197-yhkt/tiktok/tiktok-video/kitex_gen/video"
	"github.com/41197-yhkt/tiktok/tiktok-video/kitex_gen/video/douyinservice"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/constants"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var videoClient douyinservice.Client

func initVideoRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	const videoServiceName = "video"
	c, err := douyinservice.NewClient(
		videoServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(5*time.Second),              // rpc timeout
		client.WithConnectTimeout(1000*time.Millisecond),  // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

func GetVideo(ctx context.Context, req *video.GetVideoRequest) (*video.Video, error) {
	resp, err := videoClient.DouyinGetVideoMethod(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Video, nil
}

func MGetVideo(ctx context.Context, req *video.MGetVideoRequest) ([]*video.Video, error) {
	resp, err := videoClient.DouyinMGetVideoMethod(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.VideoList, nil
}
