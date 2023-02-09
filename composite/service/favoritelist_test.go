package service

import (
	"context"
	"testing"

	"github.com/41197-yhkt/tiktok/composite/kitex_gen/composite"
)

func TestFavoriteList(t *testing.T) {
	compositeClient := setupClient()
	ctx := context.Background()

	// 定义获取点赞列表请求
	var req *composite.BasicFavoriteListRequest
	// 获取点赞列表测试
	req = &composite.BasicFavoriteListRequest{UserId: 1, QueryId: 1}
	resp, err := compositeClient.BasicFavoriteListMethod(ctx, req)
	if err != nil {
		t.Error(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		t.Errorf("Error Code: %v, Error Message: %v", resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	t.Logf("Favorite List: %v", resp.VideoList)
}

func TestFavoriteListInvalidParam(t *testing.T) {
	compositeClient := setupClient()
	ctx := context.Background()

	// 定义获取点赞列表请求
	var req *composite.BasicFavoriteListRequest
	// 获取点赞列表测试
	req = &composite.BasicFavoriteListRequest{UserId: -1, QueryId: -1}
	resp, err := compositeClient.BasicFavoriteListMethod(ctx, req)
	if err != nil {
		t.Error(err)
	}
	if resp.BaseResp.StatusCode != 1001 {
		t.Errorf("Error Code: %v, Error Message: %v", resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	t.Logf("Favorite List: %v", resp.VideoList)
}
