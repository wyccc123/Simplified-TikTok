package main

import (
	"context"

	"github.com/cloudwego/simplified-tiktok/cmd/user/pack"
	"github.com/cloudwego/simplified-tiktok/cmd/user/service"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/userdemo"
	"github.com/cloudwego/simplified-tiktok/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *userdemo.RegisterRequest) (resp *userdemo.RegisterResponse, err error) {
	// TODO: Your code here...
	resp = new(userdemo.RegisterResponse)

	if len(req.UserName) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewRegisterService(ctx).Register(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *userdemo.RegisterRequest) (resp *userdemo.RegisterResponse, err error) {
	// TODO: Your code here...
	resp = new(userdemo.RegisterResponse)

	if len(req.UserName) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserId = uid
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetUserById implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserById(ctx context.Context, req *userdemo.GetUserRequest) (resp *userdemo.GetUserResponse, err error) {
	// TODO: Your code here...
	resp = new(userdemo.GetUserResponse)

	if req.UserId < 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
	}

	user, err := service.NewGetUserService(ctx).GetUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.User = user
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// Follow implements the UserServiceImpl interface.
func (s *UserServiceImpl) Follow(ctx context.Context, req *userdemo.FollowRequest) (resp *userdemo.FollowResponse, err error) {
	// TODO: Your code here...
	resp = new(userdemo.FollowResponse)

	if req.UserId < 0 || req.FollowUserId < 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewFollowService(ctx).Follow(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetFollowList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFollowList(ctx context.Context, req *userdemo.GetFollowListRequest) (resp *userdemo.GetFollowListResponse, err error) {
	// TODO: Your code here...
	resp = new(userdemo.GetFollowListResponse)

	if req.UserId < 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	user_list, err := service.NewGetFollowListService(ctx).GetFollowList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.UserList = user_list

	return resp, nil
}
