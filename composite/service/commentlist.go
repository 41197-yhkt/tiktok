package service

import (
	"context"

	"github.com/41197-yhkt/tiktok-composite/gen/dal/model"
	"github.com/41197-yhkt/tiktok-composite/kitex_gen/composite"
	"github.com/41197-yhkt/tiktok-composite/pack"
	"github.com/41197-yhkt/tiktok-composite/rpc"
	"github.com/41197-yhkt/tiktok-user/kitex_gen/user"

	"github.com/41197-yhkt/pkg/errno"
)

type CommentListService struct {
	ctx context.Context
}

func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{ctx: ctx}
}

func (s *CommentListService) CommentList(req *composite.BasicCommentListRequest) ([]*composite.Comment, error) {
	commentDatabase := q.Comment.WithContext(s.ctx)

	// 1. 从 comments 中根据 video_id 查 comments
	var comments []*model.Comment
	comments, err := commentDatabase.FindByVideoid(req.VideoId)
	if err != nil {
		return nil, errno.VideoNotExistErr
	}

	// 2. 对于每个 comments 中的 user_id users 表中查信息
	// 2.1. users 表中查信息
	authorIds := pack.AuthorIds(comments)
	authors, err := rpc.MGetUser(s.ctx, &user.CompMGetUserRequest{
		UserId:        req.UserId,
		TargetUsersId: authorIds,
	})
	if err != nil {
		return nil, err
	}

	// 2.2 relationship 表中查 user_id 和 author_id 的关注关系
	isFollows := make([]bool, len(authorIds))

	res := pack.Comments(comments, authors, isFollows)

	return res, nil
}
