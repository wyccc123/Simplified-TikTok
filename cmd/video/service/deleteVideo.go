package service

import (
	"context"

	"github.com/cloudwego/simplified-tiktok/cmd/video/dal/db"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/videodemo"
)

type DelVideoService struct {
	ctx context.Context
}

// NewDelVideoService new DelVideoService
func NewDelVideoService(ctx context.Context) *DelVideoService {
	return &DelVideoService{
		ctx: ctx,
	}
}

// DelNote video note info
func (s *DelVideoService) DelNote(req *videodemo.DeleteVideoRequest) error {
	return db.DeleteVideo(s.ctx, req.VideoId, req.UserId)
}
