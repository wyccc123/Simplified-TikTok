package rpc

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/userdemo"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/userdemo/userservice"
	constents "github.com/cloudwego/simplified-tiktok/pkg/constants"
	"github.com/cloudwego/simplified-tiktok/pkg/errno"
	"github.com/cloudwego/simplified-tiktok/pkg/middleware"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var userClient userservice.Client

func initUserRPC() {
	r, err := etcd.NewEtcdResolver([]string{constents.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constents.UserServiceName,
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
	userClient = c
}

// CreateUser create user info
func Register(ctx context.Context, req *userdemo.RegisterRequest) (int64, error) {
	resp, err := userClient.Register(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp.UserId, nil
}

func Login(ctx context.Context, req *userdemo.RegisterRequest) (int64, error) {
	resp, err := userClient.Register(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp.UserId, nil
}

func GetUserById(ctx context.Context, req *userdemo.GetUserRequest) (*userdemo.User, error) {
	resp, err := userClient.GetUserById(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.User, nil
}

func Follow(ctx context.Context, req *userdemo.FollowRequest) error {
	resp, err := userClient.Follow(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func GetFollowList(ctx context.Context, req *userdemo.GetFollowListRequest) ([]*userdemo.User, error) {
	resp, err := userClient.GetFollowList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.UserList, nil
}
