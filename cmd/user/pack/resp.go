package pack

import (
	"errors"
	"time"

	"github.com/cloudwego/simplified-tiktok/pkg/errno"

	"github.com/cloudwego/simplified-tiktok/kitex_gen/userdemo"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *userdemo.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *userdemo.BaseResp {
	return &userdemo.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
}
