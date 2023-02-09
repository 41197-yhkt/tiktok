// test
package handler

import (
	"context"
	"net/http"
	douyin "tiktok-gateway/internal/model"
	"tiktok-gateway/internal/rpc"
	"tiktok-gateway/kitex_gen/composite"

	//"github.com/41197-yhkt/pkg/constants"
	"github.com/41197-yhkt/pkg/errno"
	//"github.com/hertz-contrib/jwt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// DouyinFeedMethod .
// @router /douyin/feed [GET]
// TODO: fix feed idl
func testDouyinFeedMethod(ctx context.Context, c *app.RequestContext) {
	hlog.Info("in feed")
	var err error
	var req douyin.DouyinFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	uid := getUserIdFromJWT(ctx, c)

	errNo, videosRPC, nextTime := rpc.FeedMethod(ctx, &composite.BasicFeedRequest{
		UserId:      uid,
		LastestTime: req.LastestTime,
	})

	if errNo != *errno.Success {
		//TODO: 修复sendResponse
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

func DouyinFavoriteActionMethodTest(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinFavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// 获取uid
	// uid, errNo := getUserIdFromJWT(ctx, c)
	// if errNo != *errno.Success {
	// 	SendResponse(c, errNo)
	// 	return
	// }
	uid := getUserIdFromJWT(ctx, c)
	hlog.DefaultLogger().Info("user_id=", uid)

	// RPC调用返回true
	// errNo = rpc.FavoriteAction(context.Background(), &composite.BasicFavoriteActionRequest{
	// 	VideoId:    req.VideoID,
	// 	ActionType: req.ActionType,
	// 	UserId:     uid,
	// })

	errNo := *errno.Success

	if errNo != *errno.Success {
		SendResponse(c, errNo)
		return
	}

	SendResponse(c, errNo)
}

func testDouyinFavoriteListMethod(ctx context.Context, c *app.RequestContext) {
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

	msg := "get success"
	resp := douyin.DouyinFavoriteListResponse{
		VideoList: videosHTTP,
		BaseResp: &douyin.BaseResp{
			StatusCode: 0,
			StatusMsg:  &msg,
		},
	}

	c.JSON(http.StatusOK, resp)
}

func testDouyinCommentListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinCommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err, commentsRPC := rpc.CommentList(ctx, &composite.BasicCommentListRequest{
		VideoId: req.VideoID,
	})

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
