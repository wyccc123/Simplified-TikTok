package service

import (
	"context"

	"github.com/cloudwego/simplified-tiktok/cmd/user/pack"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/userdemo"

	"github.com/cloudwego/simplified-tiktok/cmd/user/dal/db"
)

type GetUserService struct {
	ctx context.Context
}

// NewGetUserService new GetUserService
func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

// GetUser get user info by userID
func (s *GetUserService) GetUser(req *userdemo.GetUserRequest) (*userdemo.User, error) {
	modelUser, err := db.GetUserByID(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	user := pack.User(modelUser)
	return user, nil
}
