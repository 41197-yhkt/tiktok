package service

import (
	"context"
	"log"

	"github.com/41197-yhkt/tiktok/composite/gen/dal/model"
	"github.com/41197-yhkt/tiktok/composite/kitex_gen/composite"
	"github.com/41197-yhkt/tiktok/composite/pack"
	"github.com/41197-yhkt/tiktok/composite/rpc"

	"github.com/41197-yhkt/tiktok/pkg/errno"
	"github.com/41197-yhkt/tiktok/user/kitex_gen/user"
)

type CommentActionService struct {
	ctx context.Context
}

func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{ctx: ctx}
}

func (s *CommentActionService) CommentAction(req *composite.BasicCommentActionRequest) (*composite.Comment, error) {
	commentDatabase := q.Comment.WithContext(s.ctx)

	// 1. 判断 ActionType
	if req.ActionType == 1 {
		// 2.1. 创建 Comment 数据模型，插入 comments 表中
		commentData := &model.Comment{
			UserId:  req.UserId,
			VideoId: req.VideoId,
			Content: *req.CommentText,
		}

		err := commentDatabase.Create(commentData)
		if err != nil {
			return nil, err
		}

		// 2.2. 从 comments 拿到最后一个插入的记录
		lastComment, err := commentDatabase.Last()
		if err != nil {
			return nil, err
		}

		// 2.3. users 表中查评论作者的信息
		author, err := rpc.GetUser(s.ctx, &user.CompGetUserRequest{UserId: req.UserId, TargetUserId: lastComment.UserId})
		if err != nil {
			return nil, err
		}

		// 3. 封装
		// 自己不能关注自己，所以 isFollow 固定为 false
		resInfo := pack.Comment(lastComment, author, false)
		return resInfo, nil
	} else {
		// 软删除
		comment := new(model.Comment)
		comment.ID = uint(*req.CommentId)
		resInfo, err := commentDatabase.Delete(comment)

		// 硬删除
		// err := commentDatabase.DeleteById(*req.CommentId)

		if err != nil {
			return nil, errno.CommentNotExistErr
		}
		log.Println(resInfo)
		return nil, err
	}
}
