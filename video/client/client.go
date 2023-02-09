package main

import (
	"context"
	"fmt"

	"github.com/41197-yhkt/tiktok-video/kitex_gen/video"
	"github.com/41197-yhkt/tiktok-video/kitex_gen/video/douyinservice"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var VideoClient douyinservice.Client

// 初始化 client
func initUserFavoriteRpc() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}

	c, err := douyinservice.NewClient(
		"video",
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	VideoClient = c
}

// client 模板
func main() {
	initUserFavoriteRpc()
	// 自己提供 ctx
	ctx := context.Background()

	// 调用 DouyinPublishListActionMethod 方法
	// 定义请求
	req := &video.DouyinPublishListRequest{UserId: 1}
	resp, err := VideoClient.DouyinPublishListMethod(ctx, req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	// 调用 BasicFavoriteActionMethod 方法

	// req := &video.DouyinPublishActionRequest{Author: 1, Title: "空白test", Data: []byte("123456")}
	// resp, err := VideoClient.DouyinPublishActionMethod(ctx, req)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Call DouyinPublishActionMethod: ", resp)

}
