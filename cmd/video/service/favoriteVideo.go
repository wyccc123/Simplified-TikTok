package service

import (
	"context"

	"github.com/cloudwego/simplified-tiktok/cmd/video/dal/db"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/videodemo"
	"github.com/cloudwego/simplified-tiktok/pkg/errno"
)

type FavoriteVideoService struct {
	ctx context.Context
}

// NewFavoriteVideoService new FavoriteVideoService
func NewFavoriteVideoService(ctx context.Context) *FavoriteVideoService {
	return &FavoriteVideoService{ctx: ctx}
}

func (s *FavoriteVideoService) FavoriteVideo(req *videodemo.FavoriteVideoRequest) error {
	// 1-点赞
	if req.IsFavorite == true {
		return db.FavoriteVideo(s.ctx, req.UserId, req.VideoId)
	}
	// 2-取消点赞
	if req.IsFavorite == false {
		return db.DisFavorite(s.ctx, req.UserId, req.VideoId)
	}
	return errno.ParamErr
}
