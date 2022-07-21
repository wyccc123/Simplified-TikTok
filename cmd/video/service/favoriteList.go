package service

import (
	"context"

	"github.com/cloudwego/simplified-tiktok/cmd/video/dal/db"
	"github.com/cloudwego/simplified-tiktok/cmd/video/pack"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/videodemo"
)

type FavoriteListService struct {
	ctx context.Context
}

// NewFavoriteListService creates a new FavoriteListService
func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{
		ctx: ctx,
	}
}

// FavoriteList returns a Favorite List
func (s *FavoriteListService) FavoriteList(req *videodemo.GetUserFavoriteRequest) ([]*videodemo.Video, error) {
	favorite_video_list, err := db.FavoriteList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	videos := pack.Videos(favorite_video_list)
	return videos, nil
}
