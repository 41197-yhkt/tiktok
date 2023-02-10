package handler

import (
	"bufio"
	"context"
	"log"
	"mime/multipart"

	"github.com/41197-yhkt/tiktok/gateway/biz/model/douyin"
	"github.com/41197-yhkt/tiktok/gateway/kitex_gen/video"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"

	"time"

	"github.com/41197-yhkt/tiktok/gateway/kitex_gen/video/douyinservice"
)

type DouyinPublishVideoRequest struct {
	Data  *multipart.FileHeader `form:"data"`
	Title string                `form:"title"`
	Token string                `form:"token" json:"token" query:"token"`
}

func DouyinPublishActionMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req DouyinPublishVideoRequest

	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	uid := getUserIdFromJWT(ctx, c)

	hlog.Info("title = ", req.Title, " uid=", uid)
	hlog.Info("data = ", req.Data.Filename)

	// 将文件字节化
	dataFile, err := req.Data.Open()
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
	}
	bytes := make([]byte, req.Data.Size)
	buffer := bufio.NewReader(dataFile)
	_, err = buffer.Read(bytes)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
	}

	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		return
	}

	client, err := douyinservice.NewClient("video", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}

	reqRPC := video.DouyinPublishActionRequest{
		Author: uid,
		Data:   bytes,
		Title:  req.Title,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	respRPC, err := client.DouyinPublishActionMethod(ctx, &reqRPC)
	cancel()
	if err != nil {
		log.Fatal(err)
	}

	//该函数的rpc和http的response一样，就不做转换了
	c.JSON(consts.StatusOK, respRPC)
}

func DouyinPublishListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinPublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	hlog.Info("uid=", req.UserID)

	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		return
	}

	client, err := douyinservice.NewClient("video", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	reqRPC := video.DouyinPublishListRequest{
		UserId: req.UserID,
	}
	resp, err := client.DouyinPublishListMethod(ctx, &reqRPC)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(consts.StatusOK, resp)
}

// 用于测试
func DouyinPublishActionMethodTest(ctx context.Context, c *app.RequestContext) {
	var err error
	var req DouyinPublishVideoRequest

	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	uid := getUserIdFromJWT(ctx, c)

	hlog.Info("title = ", req.Title, " uid=", uid)
	hlog.Info("data = ", req.Data.Filename)

	// 将文件字节化
	dataFile, err := req.Data.Open()
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
	}
	bytes := make([]byte, req.Data.Size)
	buffer := bufio.NewReader(dataFile)
	_, err = buffer.Read(bytes)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
	}

	hlog.Info("dataBytes=", bytes)

	//该函数的rpc和http的response一样，就不做转换了
	//c.JSON(consts.StatusOK, respRPC)
}

// 用于测试
func DouyinPublishListMethodTest(ctx context.Context, c *app.RequestContext) {
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
