package service

import (
	"context"

	"github.com/cloudwego/simplified-tiktok/cmd/comment/dal/db"
	"github.com/cloudwego/simplified-tiktok/cmd/comment/pack"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/commentdemo"
)

type GetVideoCommentService struct {
	ctx context.Context
}

func NewGetVideoCommetService(ctx context.Context) *GetVideoCommentService {
	return &GetVideoCommentService{ctx: ctx}
}

func (s GetVideoCommentService) GetVideoComment(req *commentdemo.GetVideoCommentRequest) ([]*commentdemo.Comment, error) {
	comments, err := db.GetVideoComments(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}

	comment_list := pack.Comments(comments)
	return comment_list, nil
}
