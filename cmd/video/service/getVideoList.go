package service

import (
	"context"
	"time"

	"github.com/cloudwego/simplified-tiktok/cmd/video/dal/db"
	"github.com/cloudwego/simplified-tiktok/cmd/video/pack"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/videodemo"
	constants "github.com/cloudwego/simplified-tiktok/pkg/constants"
)

type GetVideoListService struct {
	ctx context.Context
}

// NewGetVideoListService new GetVideoListService
func NewGetVideoListService(ctx context.Context) *GetVideoListService {
	return &GetVideoListService{ctx: ctx}
}

func (s *GetVideoListService) GetVideoList(req *videodemo.GetVideoListRequst) ([]*videodemo.Video, int64, error) {
	video_list, err := db.GetVideoList(s.ctx, constants.DefaultLimit, req.LatestTime)
	if err != nil {
		return nil, 0, err
	}

	if len(video_list) == 0 {
		next_time := time.Now().UnixMilli()
		return nil, next_time, nil
	}

	next_time := video_list[len(video_list)-1].UpdatedAt.UnixMilli()
	videos := pack.Videos(video_list)
	return videos, next_time, nil
}
