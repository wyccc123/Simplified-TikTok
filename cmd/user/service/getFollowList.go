package service

import (
	"context"

	"github.com/cloudwego/simplified-tiktok/cmd/user/dal/db"
	"github.com/cloudwego/simplified-tiktok/cmd/user/pack"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/userdemo"
)

type GetFollowListService struct {
	ctx context.Context
}

// NewGetFollowListService creates a new GetFollowListService
func NewGetFollowListService(ctx context.Context) *GetFollowListService {
	return &GetFollowListService{
		ctx: ctx,
	}
}

// FollowingList returns the following lists
func (s *GetFollowListService) GetFollowList(req *userdemo.GetFollowListRequest) ([]*userdemo.User, error) {
	follow_list, err := db.GetFollowList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	user_list := pack.Users(follow_list)
	return user_list, nil
}
