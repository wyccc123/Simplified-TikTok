package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/cloudwego/simplified-tiktok/cmd/user/dal/db"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/userdemo"
	"github.com/cloudwego/simplified-tiktok/pkg/errno"
)

type RegisterService struct {
	ctx context.Context
}

// NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// RegisterService create user info.
func (s *RegisterService) Register(req *userdemo.RegisterRequest) error {
	users, err := db.QueryUser(s.ctx, req.UserName)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	return db.CreateUser(s.ctx, []*db.User{{
		UserName: req.UserName,
		Password: passWord,
	}})
}
