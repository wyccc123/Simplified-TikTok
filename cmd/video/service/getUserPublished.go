package service

import (
	"context"

	"github.com/cloudwego/simplified-tiktok/cmd/video/dal/db"
	"github.com/cloudwego/simplified-tiktok/cmd/video/pack"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/videodemo"
)

type GetUserPublishedService struct {
	ctx context.Context
}

// NewGetUserPublishedService new GetUserPublishedService
func NewGetUserPublishedService(ctx context.Context) *GetUserPublishedService {
	return &GetUserPublishedService{ctx: ctx}
}

func (s *GetUserPublishedService) GetUserPublished(req *videodemo.GetUserPublishedRequest) ([]*videodemo.Video, error) {
	video_list, err := db.GetUserPublished(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	videos := pack.Videos(video_list)

	return videos, nil
}
