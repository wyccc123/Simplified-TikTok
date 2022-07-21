package main

import (
	"context"

	"github.com/cloudwego/simplified-tiktok/cmd/video/pack"
	"github.com/cloudwego/simplified-tiktok/cmd/video/service"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/videodemo"
	"github.com/cloudwego/simplified-tiktok/pkg/errno"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// PublishVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishVideo(ctx context.Context, req *videodemo.PublishRequest) (resp *videodemo.PublishResponse, err error) {
	// TODO: Your code here...
	resp = new(videodemo.PublishResponse)

	if req.UserId < 0 || len(req.VideoUrl) == 0 || len(req.VideoTitle) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewPublishService(ctx).Publish(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// DeleteVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) DeleteVideo(ctx context.Context, req *videodemo.DeleteVideoRequest) (resp *videodemo.DeleteVideoResponse, err error) {
	// TODO: Your code here...
	resp = new(videodemo.DeleteVideoResponse)

	if req.UserId < 0 || req.VideoId < 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewDelVideoService(ctx).DelNote(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetVideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoList(ctx context.Context, req *videodemo.GetVideoListRequst) (resp *videodemo.GetVideoListResponse, err error) {
	// TODO: Your code here...
	resp = new(videodemo.GetVideoListResponse)

	if req.UserId < 0 || req.LatestTime <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	video_list, next_time, err := service.NewGetVideoListService(ctx).GetVideoList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = video_list
	resp.NextTime = next_time
	return resp, nil
}

// GetUserPublished implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetUserPublished(ctx context.Context, req *videodemo.GetUserPublishedRequest) (resp *videodemo.GetUserPublishedResponse, err error) {
	// TODO: Your code here...
	resp = new(videodemo.GetUserPublishedResponse)

	if req.UserId < 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	video_list, err := service.NewGetUserPublishedService(ctx).GetUserPublished(req)
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = video_list
	return resp, nil
}

// FavoriteVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteVideo(ctx context.Context, req *videodemo.FavoriteVideoRequest) (resp *videodemo.FavoriteVideoResponse, err error) {
	// TODO: Your code here...
	resp = new(videodemo.FavoriteVideoResponse)

	if req.UserId < 0 || req.VideoId < 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	if req.IsFavorite != true && req.IsFavorite != false {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewFavoriteVideoService(ctx).FavoriteVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetUserFavorite implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetUserFavorite(ctx context.Context, req *videodemo.GetUserFavoriteRequest) (resp *videodemo.GetUserFavoriteResponse, err error) {
	// TODO: Your code here...
	resp = new(videodemo.GetUserFavoriteResponse)

	if req.UserId < 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	videos, err := service.NewFavoriteListService(ctx).FavoriteList(req)
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = videos
	return resp, nil
}
