package service

import (
	"context"

	"github.com/cloudwego/simplified-tiktok/cmd/comment/dal/db"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/commentdemo"
)

type GiveCommentService struct {
	ctx context.Context
}

func NewGiveCommentService(ctx context.Context) *GiveCommentService {
	return &GiveCommentService{
		ctx: ctx,
	}
}

func (s GiveCommentService) GiveComment(req *commentdemo.GiveCommentRequest) error {
	return db.NewComment(s.ctx, &db.Comment{
		UserID:  int(req.UserId),
		VideoID: int(req.VideoId),
		Content: req.Content,
	})
}
