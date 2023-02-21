// 原本
package handler

import (
	"context"
	"net/http"

	douyin "github.com/41197-yhkt/tiktok/gateway/internal/model"
	"github.com/41197-yhkt/tiktok/gateway/internal/rpc"
	"github.com/41197-yhkt/tiktok/gateway/kitex_gen/composite"

	"github.com/41197-yhkt/tiktok/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
)

// DouyinFeedMethod .
// @router /douyin/feed [GET]
// TODO: fix feed idl
func DouyinFeedMethod(ctx context.Context, c *app.RequestContext) {
	hlog.Info("in feed")
	var err error
	var req douyin.DouyinFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	tokenString := c.Query("token")
	var uid int64
	if tokenString != "" {
		mw, _ := jwt.New(&jwt.HertzJWTMiddleware{
			Key: []byte("tiktok"), //密钥
		})

		token, err := mw.ParseTokenString(tokenString)
		if err != nil {
			c.String(consts.StatusBadRequest, err.Error())
			return
		}

		claims := jwt.ExtractClaimsFromToken(token)
		uid = int64(claims["identity"].(float64))
	} else {
		uid = 0
	}
	hlog.Info("uid = ", uid)

	errNo, videosRPC, nextTime := rpc.FeedMethod(ctx, &composite.BasicFeedRequest{
		UserId:      uid,
		LastestTime: req.LastestTime,
	})

	if errNo != *errno.Success {
		hlog.Info("err: ", errNo.Msg)
		SendResponse(c, errNo)
		return
	}

	// 烦人的类型转换
	var videosHTTP []*douyin.Video

	for _, v := range videosRPC {
		videoHttp := douyin.Video{
			ID: v.Id,
			Author: &douyin.User{
				FollowerCount: v.Author.FollowerCount,
				Name:          v.Author.Name,
				ID:            v.Author.Id,
				FollowCount:   v.Author.FollowCount,
				IsFollow:      v.Author.IsFollow,
			},
			PlayURL:       v.PlayUrl,
			CoverURL:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			Title:         v.Title,
			IsFavorite:    v.IsFavorite,
		}
		videosHTTP = append(videosHTTP, &videoHttp)
	}

	resp := douyin.DouyinFeedResponse{
		VideoList: videosHTTP,
		NextTime:  nextTime,
		BaseResp: &douyin.BaseResp{
			StatusCode: 0,
		},
	}

	c.JSON(consts.StatusOK, resp)
}

// DouyinFavoriteActionMethod .
// @router /douyin/favorite/action [POST]
func DouyinFavoriteActionMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinFavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// 获取uid
	uid := getUserIdFromJWT(ctx, c)

	errNo := rpc.FavoriteAction(context.Background(), &composite.BasicFavoriteActionRequest{
		VideoId:    req.VideoID,
		ActionType: req.ActionType,
		UserId:     uid,
	})

	if errNo != *errno.Success {
		hlog.Info("in favourite action err: ", errNo.Msg)
		SendResponse(c, errNo)
		return
	}

	SendResponse(c, errNo)
}

// DouyinFavoriteListMethod .
// @router /douyin/favorite/list [GET]
func DouyinFavoriteListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinFavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	errNo, videosRPC := rpc.FavoriteList(ctx, &composite.BasicFavoriteListRequest{
		UserId:  req.UserID,
		QueryId: req.UserID,
	})
	if errNo != *errno.Success {
		hlog.Info("in favourite List err: ", errNo.Msg)
		SendResponse(c, errNo)
		return
	}

	// 烦人的类型转换
	var videosHTTP []*douyin.Video

	for _, v := range videosRPC {
		videoHttp := douyin.Video{
			ID: v.Id,
			Author: &douyin.User{
				FollowerCount: v.Author.FollowerCount,
				Name:          v.Author.Name,
				ID:            v.Author.Id,
				FollowCount:   v.Author.FollowCount,
				IsFollow:      v.Author.IsFollow,
			},
			PlayURL:       v.PlayUrl,
			CoverURL:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			Title:         v.Title,
			IsFavorite:    v.IsFavorite,
		}
		videosHTTP = append(videosHTTP, &videoHttp)
	}

	resp := douyin.DouyinFavoriteListResponse{
		VideoList: videosHTTP,
		BaseResp: &douyin.BaseResp{
			StatusCode: 0,
		},
	}

	c.JSON(http.StatusOK, resp)
}

// DouyinCommentActionMethod .
// @router /douyin/comment/action [POST]
func DouyinCommentActionMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinCommentActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	uid := getUserIdFromJWT(ctx, c)
	hlog.Info("uid = ", uid)
	errNo, rpcResp := rpc.CommentAction(context.Background(), &composite.BasicCommentActionRequest{
		VideoId:     req.VideoID,
		UserId:      uid,
		ActionType:  req.ActionType,
		CommentId:   req.CommentID,
		CommentText: req.CommentText,
	})

	if errNo != *errno.Success {
		hlog.Info("in Comment Action err: ", errNo.Msg)
		SendResponse(c, errNo)
		return
	}

	var resp douyin.DouyinCommentActionResponse
	if req.ActionType == 1 {
		resp = douyin.DouyinCommentActionResponse{
			BaseResp: &douyin.BaseResp{
				StatusCode: 0,
			},
			Comment: &douyin.Comment{
				ID:      uid,
				Content: *req.CommentText,
				User: &douyin.User{
					ID:            rpcResp.Comment.Id,
					Name:          rpcResp.Comment.User.Name,
					FollowCount:   rpcResp.Comment.User.FollowCount,
					FollowerCount: rpcResp.Comment.User.FollowerCount,
					IsFollow:      rpcResp.Comment.User.IsFollow,
				},
				CreateDate: *rpcResp.BaseResp.StatusMsg,
			},
		}
	} else {
		resp = douyin.DouyinCommentActionResponse{
			BaseResp: &douyin.BaseResp{
				StatusCode: 0,
			},
		}
	}

	c.JSON(consts.StatusOK, resp)
}

// DouyinCommentListMethod .
// @router /douyin/comment/list [GET]
func DouyinCommentListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinCommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	uid := getUserIdFromJWT(ctx, c)
	hlog.Info("uid=", uid)

	errNo, commentsRPC := rpc.CommentList(ctx, &composite.BasicCommentListRequest{
		VideoId: req.VideoID,
		UserId:  uid,
	})

	if errNo != *errno.Success {
		hlog.Info("in Comment List err: ", errNo.Msg)
		SendResponse(c, errNo)
		return
	}
	var commentsHTTP []*douyin.Comment
	// 类型转换
	for _, c := range commentsRPC {
		commentHTTP := douyin.Comment{
			ID: c.Id,
			User: &douyin.User{
				FollowerCount: c.User.FollowerCount,
				Name:          c.User.Name,
				ID:            c.User.Id,
				FollowCount:   c.User.FollowCount,
				IsFollow:      c.User.IsFollow,
			},
			Content:    c.Content,
			CreateDate: c.CreateDate,
		}
		commentsHTTP = append(commentsHTTP, &commentHTTP)
	}

	resp := douyin.DouyinCommentListResponse{
		CommentList: commentsHTTP,
		BaseResp: &douyin.BaseResp{
			StatusCode: 0,
		},
	}

	c.JSON(consts.StatusOK, resp)
}

func getUserIdFromJWT(ctx context.Context, c *app.RequestContext) int64 {
	claim := jwt.ExtractClaims(ctx, c)
	uid := int64(claim["identity"].(float64))
	return uid
}
