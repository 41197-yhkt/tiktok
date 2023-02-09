package handler

import (
	"context"
	"log"
	"tiktok-gateway/kitex_gen/video"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"

	"tiktok-gateway/kitex_gen/video/douyinservice"
	"time"
)

// DouyinPublishActionMethod .
// @router /douyin/publish/action [POST]
func DouyinPublishActionMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req video.DouyinPublishActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		return
	}

	client, err := douyinservice.NewClient("video", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	resp, err := client.DouyinPublishActionMethod(ctx, &req)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(consts.StatusOK, resp)
}

// DouyinPublishListMethod .
// @router /douyin/publish/list [GET]
func DouyinPublishListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req video.DouyinPublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		return
	}

	client, err := douyinservice.NewClient("video", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	resp, err := client.DouyinPublishListMethod(ctx, &req)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(consts.StatusOK, resp)
}
