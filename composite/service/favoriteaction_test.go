package service

import (
	"context"
	"testing"

	"github.com/41197-yhkt/tiktok-composite/kitex_gen/composite"
)

func TestFavoriteAction(t *testing.T) {
	compositeClient := setupClient()
	ctx := context.Background()

	// 定义点赞请求
	var req *composite.BasicFavoriteActionRequest
	// 点赞测试
	req = &composite.BasicFavoriteActionRequest{VideoId: 5, UserId: 1, ActionType: 1}
	resp, err := compositeClient.BasicFavoriteActionMethod(ctx, req)
	if err != nil {
		t.Error(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		t.Errorf("Error Code: %v, Error Message: %v", resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}

	// 取消点赞测试
	req = &composite.BasicFavoriteActionRequest{VideoId: 5, UserId: 1, ActionType: 2}
	resp, err = compositeClient.BasicFavoriteActionMethod(ctx, req)
	if err != nil {
		t.Error(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		t.Errorf("Error Code: %v, Error Message: %v", resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
}

func TestFavoriteActionInvalidParam(t *testing.T) {
	compositeClient := setupClient()
	ctx := context.Background()

	// 定义点赞请求
	var req *composite.BasicFavoriteActionRequest
	// 点赞测试
	req = &composite.BasicFavoriteActionRequest{VideoId: -5, UserId: -1, ActionType: 1}
	resp, err := compositeClient.BasicFavoriteActionMethod(ctx, req)
	if err != nil {
		t.Error(err)
	}
	if resp.BaseResp.StatusCode != 1001 {
		t.Errorf("Error Code: %v, Error Message: %v", resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
}
