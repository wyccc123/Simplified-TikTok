package pack

import (
	"errors"
	"time"

	"github.com/cloudwego/simplified-tiktok/pkg/errno"

	"github.com/cloudwego/simplified-tiktok/kitex_gen/videodemo"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *videodemo.BaseResp {
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

func baseResp(err errno.ErrNo) *videodemo.BaseResp {
	return &videodemo.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
}
