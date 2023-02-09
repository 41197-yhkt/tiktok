package rpc

import (
	"context"
	"tiktok-gateway/kitex_gen/composite"
	"tiktok-gateway/kitex_gen/composite/compositeservice"
	"time"

	"github.com/41197-yhkt/pkg/constants"
	"github.com/41197-yhkt/pkg/errno"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var compClient compositeservice.Client

func initCompRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := compositeservice.NewClient(
		constants.CompServiceName,
		//client.WithMiddleware(middleware.CommonMiddleware),
		//client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		// panic(err)
	}
	compClient = c
}

func FavoriteAction(ctx context.Context, req *composite.BasicFavoriteActionRequest) errno.ErrNo {
	resp, err := compClient.BasicFavoriteActionMethod(ctx, req)
	if err != nil {
		return *errno.ServerError
	}
	if resp.BaseResp.StatusCode != 0 {
		errInt := int(resp.BaseResp.StatusCode)
		return *errno.NewErrNo(errInt, *resp.BaseResp.StatusMsg)
	}
	return *errno.Success
}

func FavoriteList(ctx context.Context, req *composite.BasicFavoriteListRequest) (errno.ErrNo, []*composite.Video) {
	resp, err := compClient.BasicFavoriteListMethod(ctx, req)

	if err != nil {
		return *errno.ServerError, nil
	}
	if resp.BaseResp.StatusCode != 0 {
		errInt := int(resp.BaseResp.StatusCode)
		return *errno.NewErrNo(errInt, *resp.BaseResp.StatusMsg), nil
	}

	return *errno.Success, resp.VideoList
}

func FeedMethod(ctx context.Context, req *composite.BasicFeedRequest) (errno.ErrNo, []*composite.Video, *int64) {
	resp, err := compClient.BasicFeedMethod(ctx, req)

	if err != nil {
		return *errno.ServerError, nil, nil
	}
	if resp.BaseResp.StatusCode != 0 {
		errInt := int(resp.BaseResp.StatusCode)
		return *errno.NewErrNo(errInt, *resp.BaseResp.StatusMsg), nil, nil
	}

	return *errno.Success, resp.VideoList, resp.NextTime
}

func CommentAction(ctx context.Context, req *composite.BasicCommentActionRequest) (errno.ErrNo, *composite.BasicCommentActionResponse) {
	resp, err := compClient.BasicCommentActionMethod(ctx, req)

	if err != nil {
		return *errno.ServerError, nil
	}
	if resp.BaseResp.StatusCode != 0 {
		errInt := int(resp.BaseResp.StatusCode)
		return *errno.NewErrNo(errInt, *resp.BaseResp.StatusMsg), nil
	}
	return *errno.Success, resp
}

func CommentList(ctx context.Context, req *composite.BasicCommentListRequest) (errno.ErrNo, []*composite.Comment) {
	resp, err := compClient.BasicCommentListMethod(ctx, req)

	if err != nil {
		return *errno.ServerError, nil
	}
	if resp.BaseResp.StatusCode != 0 {
		errInt := int(resp.BaseResp.StatusCode)
		return *errno.NewErrNo(errInt, *resp.BaseResp.StatusMsg), nil
	}
	return *errno.Success, resp.CommentList
}
