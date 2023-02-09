package main

import (
	"context"

	video "github.com/41197-yhkt/tiktok-video/kitex_gen/video"
	"github.com/41197-yhkt/tiktok-video/pack"
	service "github.com/41197-yhkt/tiktok-video/service"

	errno "github.com/41197-yhkt/pkg/errno"
)

// DouyinServiceImpl implements the last service interface defined in the IDL.
type DouyinServiceImpl struct{}

// DouyinPublishActionMethod implements the DouyinServiceImpl interface.
func (s *DouyinServiceImpl) DouyinPublishActionMethod(ctx context.Context, req *video.DouyinPublishActionRequest) (resp *video.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	resp = new(video.DouyinPublishActionResponse)

	//检验参数规范性
	if req.Author < 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.InvalidParams)
		return resp, nil
	}

	//调用服务
	err = service.NewPublishActionService(ctx).PublishAction(req)

	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// DouyinPublishListMethod implements the DouyinServiceImpl interface.
func (s *DouyinServiceImpl) DouyinPublishListMethod(ctx context.Context, req *video.DouyinPublishListRequest) (resp *video.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	resp = new(video.DouyinPublishListResponse)

	//检验参数规范性
	if req.UserId < 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.InvalidParams)
		return resp, nil
	}

	//调用服务
	resp.VideoList, err = service.NewPublishListService(ctx).PublishList(req)

	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// DouyinGetVideoMethod implements the DouyinServiceImpl interface.
func (s *DouyinServiceImpl) DouyinGetVideoMethod(ctx context.Context, req *video.GetVideoRequest) (resp *video.GetVideoResponse, err error) {
	// TODO: Your code here...
	resp = new(video.GetVideoResponse)

	//调用服务
	resp.Video, err = service.NewGetVideoService(ctx).GetVideo(req)

	return resp, err
}

// DouyinMGetVideoMethod implements the DouyinServiceImpl interface.
func (s *DouyinServiceImpl) DouyinMGetVideoMethod(ctx context.Context, req *video.MGetVideoRequest) (resp *video.MGetVideoResponse, err error) {
	// TODO: Your code here...
	resp = new(video.MGetVideoResponse)
	var res *video.Video
	for _, vid := range req.TargetVideosId {
		r := new(video.GetVideoRequest)
		r.UserId = req.UserId
		r.TargetVideoId = vid

		res, err = service.NewGetVideoService(ctx).GetVideo(r)
		resp.VideoList = append(resp.VideoList, res)
	}
	return resp, err
}
