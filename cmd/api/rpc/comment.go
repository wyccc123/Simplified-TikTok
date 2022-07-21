package rpc

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/commentdemo"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/commentdemo/commentservice"
	constents "github.com/cloudwego/simplified-tiktok/pkg/constants"
	"github.com/cloudwego/simplified-tiktok/pkg/errno"
	"github.com/cloudwego/simplified-tiktok/pkg/middleware"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var commentClient commentservice.Client

func initCommentRPC() {
	r, err := etcd.NewEtcdResolver([]string{constents.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := commentservice.NewClient(
		constents.CommentServiceName,
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
	commentClient = c
}

func GiveComment(ctx context.Context, req *commentdemo.GiveCommentRequest) error {
	resp, err := commentClient.GiveComment(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return nil
}

func DeleteComment(ctx context.Context, req *commentdemo.DeleteCommentRequest) error {
	resp, err := commentClient.DeleteComment(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return nil
}

func GetVideoComment(ctx context.Context, req *commentdemo.GetVideoCommentRequest) ([]*commentdemo.Comment, error) {
	resp, err := commentClient.GetVideoComment(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp.CommentList, nil
}
