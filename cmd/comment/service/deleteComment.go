package service

import (
	"context"

	"github.com/cloudwego/simplified-tiktok/cmd/comment/dal/db"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/commentdemo"
)

type DeleteCommentService struct {
	ctx context.Context
}

func NewDeleteCommentService(ctx context.Context) *DeleteCommentService {
	return &DeleteCommentService{
		ctx: ctx,
	}
}

func (s DeleteCommentService) DeleteComment(req *commentdemo.DeleteCommentRequest) error {
	return db.DelComment(s.ctx, req.CommentId)
}
