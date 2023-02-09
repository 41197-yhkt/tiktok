package service

import (
	"context"
	"testing"

	"github.com/41197-yhkt/tiktok-composite/kitex_gen/composite"
)

func TestCommentList(t *testing.T) {
	compositeClient := setupClient()
	ctx := context.Background()

	// 定义评论请求
	var req *composite.BasicCommentListRequest
	// 获取评论列表测试
	req = &composite.BasicCommentListRequest{
		VideoId: 20086,
	}
	resp, err := compositeClient.BasicCommentListMethod(ctx, req)
	if err != nil {
		t.Error(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		t.Errorf("Error Code: %v, Error Message: %v", resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	t.Log(resp.CommentList)
}

func TestCommentListInvalidParam(t *testing.T) {
	compositeClient := setupClient()
	ctx := context.Background()

	// 定义评论请求
	var req *composite.BasicCommentListRequest
	// 获取评论列表测试
	req = &composite.BasicCommentListRequest{
		VideoId: -1,
	}
	resp, err := compositeClient.BasicCommentListMethod(ctx, req)
	if err != nil {
		t.Error(err)
	}
	if resp.BaseResp.StatusCode != 1001 {
		t.Errorf("Error Code: %v, Error Message: %v", resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	t.Log(resp.CommentList)
}
