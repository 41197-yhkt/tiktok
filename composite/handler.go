package main

import (
	"context"

	"github.com/41197-yhkt/tiktok/composite/kitex_gen/composite"
	"github.com/41197-yhkt/tiktok/composite/service"

	"github.com/41197-yhkt/tiktok/composite/pack"

	"github.com/41197-yhkt/tiktok/pkg/errno"
)

// CompositeServiceImpl implements the last service interface defined in the IDL.
type CompositeServiceImpl struct{}

// BasicFavoriteActionMethod implements the CompositeServiceImpl interface.
func (s *CompositeServiceImpl) BasicFavoriteActionMethod(ctx context.Context, req *composite.BasicFavoriteActionRequest) (resp *composite.BasicFavoriteActionResponse, err error) {
	resp = new(composite.BasicFavoriteActionResponse)

	// 检验参数是否符合规范
	if req.VideoId <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.InvalidParams)
		return resp, nil
	}

	err = service.NewFavoriteActionService(ctx).FavoriteAction(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// BasicFavoriteListMethod implements the CompositeServiceImpl interface.
func (s *CompositeServiceImpl) BasicFavoriteListMethod(ctx context.Context, req *composite.BasicFavoriteListRequest) (resp *composite.BasicFavoriteListResponse, err error) {
	resp = new(composite.BasicFavoriteListResponse)

	// 检验参数是否符合规范
	if req.UserId <= 0 || req.QueryId <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.InvalidParams)
		return resp, nil
	}

	resp.VideoList, err = service.NewFavoriteListService(ctx).FavoriteList(req)
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// BasicFeedMethod implements the CompositeServiceImpl interface.
func (s *CompositeServiceImpl) BasicFeedMethod(ctx context.Context, req *composite.BasicFeedRequest) (resp *composite.BasicFeedResponse, err error) {
	resp = new(composite.BasicFeedResponse)

	videoList, nextTime, err := service.NewFeedService(ctx).Feed(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return nil, err
	}

	resp.VideoList = videoList
	resp.NextTime = &nextTime
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// BasicCommentActionMethod implements the CompositeServiceImpl interface.
func (s *CompositeServiceImpl) BasicCommentActionMethod(ctx context.Context, req *composite.BasicCommentActionRequest) (resp *composite.BasicCommentActionResponse, err error) {
	resp = new(composite.BasicCommentActionResponse)

	// 检验参数是否符合规范
	if (req.ActionType == 1 && (req.VideoId <= 0 || req.UserId <= 0)) || (req.ActionType == 2 && (*req.CommentId <= 0)) {
		resp.BaseResp = pack.BuildBaseResp(errno.InvalidParams)
		return resp, nil
	}

	resp.Comment, err = service.NewCommentActionService(ctx).CommentAction(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// BasicCommentListMethod implements the CompositeServiceImpl interface.
func (s *CompositeServiceImpl) BasicCommentListMethod(ctx context.Context, req *composite.BasicCommentListRequest) (resp *composite.BasicCommentListResponse, err error) {
	resp = new(composite.BasicCommentListResponse)

	if req.VideoId <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.InvalidParams)
		return resp, nil
	}

	resp.CommentList, err = service.NewCommentListService(ctx).CommentList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
