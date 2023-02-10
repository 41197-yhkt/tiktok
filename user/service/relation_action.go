package service

import (
	"context"

	"github.com/41197-yhkt/tiktok/pkg/errno"
	"github.com/41197-yhkt/tiktok/user/kitex_gen/user"
	"github.com/41197-yhkt/tiktok/user/util"
	"github.com/opentracing/opentracing-go"
)

func RelationAction(ctx context.Context, req *user.RelationActionRequest) (resp *user.RelationActionResponse, err error) {
	resp = user.NewRelationActionResponse()
	span, ctx := opentracing.StartSpanFromContext(ctx, "RelationAction")
	defer span.Finish()
	action_type := req.ActionType
	userId := req.UserId
	toUserId := req.ToUserId

	if action_type == 1 {
		var userFlollowReq = user.NewUserFollowRequest()
		userFlollowReq.FollowFrom = userId
		userFlollowReq.FollowTo = toUserId
		resp1, err := UserFollow(ctx, userFlollowReq)
		if err != nil {
			resp.BaseResp = util.PackBaseResp(err)
			return resp, nil
		}
		resp.BaseResp = resp1.BaseResp
		return resp, nil
	} else if action_type == 2 {
		var userUnFollowReq = user.NewUserUnfollowRequest()
		userUnFollowReq.FollowFrom = userId
		userUnFollowReq.FollowTo = toUserId
		resp2, err := UserUnfollow(ctx, userUnFollowReq)
		if err != nil {
			resp.BaseResp = util.PackBaseResp(err)
			return resp, nil
		}
		resp.BaseResp = resp2.BaseResp
		return resp, nil
	} else {
		return resp, errno.InvalidParams
	}

}
