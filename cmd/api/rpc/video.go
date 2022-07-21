package rpc

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/videodemo"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/videodemo/videoservice"
	constents "github.com/cloudwego/simplified-tiktok/pkg/constants"
	"github.com/cloudwego/simplified-tiktok/pkg/errno"
	"github.com/cloudwego/simplified-tiktok/pkg/middleware"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var videoClient videoservice.Client

func initVideoRPC() {
	r, err := etcd.NewEtcdResolver([]string{constents.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := videoservice.NewClient(
		constents.VideoServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

func Publish(ctx context.Context, req *videodemo.PublishRequest) error {
	resp, err := videoClient.PublishVideo(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return nil
}

func DeleteVideo(ctx context.Context, req *videodemo.DeleteVideoRequest) error {
	resp, err := videoClient.DeleteVideo(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return nil
}

func GetVideoList(ctx context.Context, req *videodemo.GetVideoListRequst) ([]*videodemo.Video, int64, error) {
	resp, err := videoClient.GetVideoList(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp.VideoList, resp.NextTime, nil
}

func GetUserPublished(ctx context.Context, req *videodemo.GetUserPublishedRequest) ([]*videodemo.Video, error) {
	resp, err := videoClient.GetUserPublished(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp.VideoList, nil
}

func FavoriteVideo(ctx context.Context, req *videodemo.FavoriteVideoRequest) error {
	resp, err := videoClient.FavoriteVideo(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return nil
}

func GetUserFavorite(ctx context.Context, req *videodemo.GetUserFavoriteRequest) ([]*videodemo.Video, error) {
	resp, err := videoClient.GetUserFavorite(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp.VideoList, nil
}
