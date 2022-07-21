package service

import (
	"context"

	"github.com/cloudwego/simplified-tiktok/cmd/user/dal/db"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/userdemo"
	"github.com/cloudwego/simplified-tiktok/pkg/errno"
)

type FollowService struct {
	ctx context.Context
}

// NewFollowService new FollowService
func NewFollowService(ctx context.Context) *FollowService {
	return &FollowService{ctx: ctx}
}

func (s *FollowService) Follow(req *userdemo.FollowRequest) error {
	// 1-关注
	if req.IsFollow == true {
		return db.NewFollow(s.ctx, req.UserId, req.FollowUserId)
	}
	// 2-取消关注
	if req.IsFollow == false {
		return db.DisFollow(s.ctx, req.UserId, req.FollowUserId)
	}
	return errno.ParamErr
}
