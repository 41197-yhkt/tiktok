package service

import (
	"context"
	"testing"

	"github.com/41197-yhkt/tiktok/tiktok-composite/kitex_gen/composite"
	"github.com/41197-yhkt/tiktok/tiktok-composite/kitex_gen/composite/compositeservice"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/constants"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func setupClient() compositeservice.Client {
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
	return c
}

func TestCommentAction(t *testing.T) {
	compositeClient := setupClient()
	ctx := context.Background()

	// 定义评论请求
	var req *composite.BasicCommentActionRequest
	// 评论测试
	commentText := "你好哈哈"
	req = &composite.BasicCommentActionRequest{
		UserId:      10086,
		VideoId:     20086,
		ActionType:  1,
		CommentText: &commentText,
	}
	resp, err := compositeClient.BasicCommentActionMethod(ctx, req)
	if err != nil {
		t.Error(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		t.Errorf("Error Code: %v, Error Message: %v", resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	t.Log(resp.Comment)

	// 删除评论测试
	req = &composite.BasicCommentActionRequest{
		ActionType: 2,
		CommentId:  &resp.Comment.Id,
	}
	resp, err = compositeClient.BasicCommentActionMethod(ctx, req)
	if err != nil {
		t.Error(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		t.Errorf("Error Code: %v, Error Message: %v", resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	t.Log(resp.Comment)
}

func TestCommentActionInvalidParam(t *testing.T) {
	compositeClient := setupClient()
	ctx := context.Background()

	// 定义评论请求
	var req *composite.BasicCommentActionRequest
	// 评论测试
	commentText := "你好哈哈"
	req = &composite.BasicCommentActionRequest{
		UserId:      -10086,
		VideoId:     -20086,
		ActionType:  1,
		CommentText: &commentText,
	}
	resp, err := compositeClient.BasicCommentActionMethod(ctx, req)
	if err != nil {
		t.Error(err)
	}
	if resp.BaseResp.StatusCode != 1001 {
		t.Errorf("Error Code: %v, Error Message: %v", resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	t.Log(resp.Comment)

}
