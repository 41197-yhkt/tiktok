package main

import (
	"context"
	"fmt"

	"github.com/41197-yhkt/tiktok-composite/kitex_gen/composite"
	"github.com/41197-yhkt/tiktok-composite/kitex_gen/composite/compositeservice"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/constants"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var compositeClient compositeservice.Client

// 初始化 client
func initUserFavoriteRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := compositeservice.NewClient(
		"composite",
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	compositeClient = c
}

// client 模板
func main() {
	initUserFavoriteRpc()
	// 自己提供 ctx
	ctx := context.Background()

	// 调用 BasicFavoriteActionMethod 方法
	// 定义请求
	// req := &composite.BasicFavoriteActionRequest{VideoId: 1, UserId: 1}
	// resp, err := compositeClient.BasicFavoriteActionMethod(ctx, req)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(resp)

	// 调用 BasicFavoriteListMethod 方法
	req := &composite.BasicFavoriteListRequest{UserId: 2, QueryId: 1}
	resp, err := compositeClient.BasicFavoriteListMethod(ctx, req)
	if err != nil {
		panic(err)
	}
	fmt.Println("Call BasicFavoriteListMethod: ", resp)

	// 调用 CommentActionMethod 方法增加评论
	// commentText := "你好哈哈"
	// req := &composite.BasicCommentActionRequest{
	// 	UserId:      1,
	// 	VideoId:     2,
	// 	ActionType:  1,
	// 	CommentText: &commentText,
	// }
	// resp, err := compositeClient.BasicCommentActionMethod(ctx, req)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Call BadicCommentActionMethod to add comment: ", resp)

	// 调用 CommentMethod 方法删除评论
	// var deleteId int64 = 1
	// req := &composite.BasicCommentActionRequest{
	// 	ActionType: 2,
	// 	CommentId:  &deleteId,
	// }
	// resp, err := compositeClient.BasicCommentActionMethod(ctx, req)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Call BadicCommentActionMethod to delete comment: ", resp)

	// 调用 CommentList 获取评论列表
	// req := &composite.BasicCommentListRequest{VideoId: 1}
	// resp, err := compositeClient.BasicCommentListMethod(ctx, req)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Call BasicCommentListMethod: ", resp)

	// 调用 Feed 获取视屏列表
	// req := &composite.BasicFeedRequest{}
	// resp, err := compositeClient.BasicFeedMethod(ctx, req)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(resp)
}
