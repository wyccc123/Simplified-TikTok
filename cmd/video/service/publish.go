package service

import (
	"context"

	"github.com/cloudwego/simplified-tiktok/cmd/video/dal/db"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/videodemo"
)

type PublishService struct {
	ctx context.Context
}

// NewPublishService new PublishService
func NewPublishService(ctx context.Context) *PublishService {
	return &PublishService{ctx: ctx}
}

func (s *PublishService) Publish(req *videodemo.PublishRequest) error {
	return db.PublishVideo(s.ctx, &db.Video{
		AuthorID:      int(req.UserId),
		PlayUrl:       req.VideoUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         req.VideoTitle,
	})
}
