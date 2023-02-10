package handler

import (
	"context"
	"log"
	"time"

	// "github.com/41197-yhkt/tiktok/gateway/kitex_gen/relation"
	// "github.com/41197-yhkt/tiktok/gateway/kitex_gen/relation/relationservice"
	"github.com/41197-yhkt/tiktok/gateway/kitex_gen/relation"
	"github.com/41197-yhkt/tiktok/gateway/kitex_gen/relation/relationservice"
	"github.com/41197-yhkt/tiktok/gateway/kitex_gen/user"
	"github.com/41197-yhkt/tiktok/gateway/kitex_gen/user/userservice"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// DouyinUserRegisterMethod .
// @router /relation/user/register [POST]
func DouyinUserRegisterMethod(ctx context.Context, c *app.RequestContext) {
	//return
	var err error
	var req user.UserRegisterRequest
	err = c.BindAndValidate(&req)
	log.Print(req)
	if err != nil {
		log.Fatal("Bind ERROR")
	}

	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	client, err := userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	_, err = client.UserRegister(ctx, &req)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
}

// DouyinUserLoginMethod .
// @router /relation/user/login [POST]
func DouyinUserLoginMethod(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	return &user.UserLoginResponse{
		BaseResp: nil,
		UserId:   1,
	}, nil
	// var err error
	// var req user.UserLoginRequest
	// err = c.BindAndValidate(&req)
	// if err != nil {
	// 	return nil, err
	// }
	// r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	// if err != nil {
	// 	return nil, err
	// }

	// client, err := userservice.NewClient("user", client.WithResolver(r))
	// if err != nil {
	// 	return nil, err
	// }
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	// resp, err := client.UserLogin(ctx, &req)
	// cancel()
	// if err != nil {
	// 	return nil, err
	// }

	// return resp, nil
}

// DouyinUserMethod .
// @router /relation/user [GET]
func DouyinUserMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserInfoRequest
	err = c.BindAndValidate(&req)
	log.Print(req)
	if err != nil {
		log.Fatal("Bind ERROR")
	}

	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	client, err := userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	resp, err := client.UserInfo(ctx, &req)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(consts.StatusOK, resp)
}

// DouyinRelationActionMethod .
// @router /relation/relation/action [POST]
func DouyinRelationActionMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.DouyinRelationActionRequest
	err = c.BindAndValidate(&req)
	log.Print(req)
	if err != nil {
		log.Fatal("Bind ERROR")
	}
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	client, err := relationservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	resp, err := client.DouyinRelationActionMethod(ctx, &req)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(consts.StatusOK, resp)
}

// DouyinRelationFollowListMethod .
// @router /relation/relation/follow/list [GET]
func DouyinRelationFollowListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.FollowListRequest
	err = c.BindAndValidate(&req)
	log.Print(req)
	if err != nil {
		log.Fatal("Bind ERROR")
	}
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	client, err := userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	resp, err := client.GetFollowList(ctx, &req)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(consts.StatusOK, resp)
}

// DouyinRelationFollowerListMethod .
// @router /relation/relation/follower/list [GET]
func DouyinRelationFollowerListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.FollowerListRequest
	err = c.BindAndValidate(&req)
	log.Print(req)
	if err != nil {
		log.Fatal("Bind ERROR")
	}
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	client, err := userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	resp, err := client.GetFollowerList(ctx, &req)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(consts.StatusOK, resp)
}

// DouyinRelationFriendListMethod .
// @router /relation/relation/friend/list [GET]
func DouyinRelationFriendListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.FriendListRequest
	err = c.BindAndValidate(&req)
	log.Print(req)
	if err != nil {
		log.Fatal("Bind ERROR")
	}
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	client, err := userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	resp, err := client.GetFriendList(ctx, &req)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(consts.StatusOK, resp)
}
