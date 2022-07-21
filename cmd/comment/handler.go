package main

import (
	"context"

	"github.com/cloudwego/simplified-tiktok/cmd/comment/pack"
	"github.com/cloudwego/simplified-tiktok/cmd/comment/service"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/commentdemo"
	"github.com/cloudwego/simplified-tiktok/pkg/errno"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// GiveComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) GiveComment(ctx context.Context, req *commentdemo.GiveCommentRequest) (resp *commentdemo.GiveCommentResponse, err error) {
	// TODO: Your code here...
	resp = new(commentdemo.GiveCommentResponse)

	if req.UserId < 0 || req.VideoId < 0 || len(req.Content) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewGiveCommentService(ctx).GiveComment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// DeleteComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) DeleteComment(ctx context.Context, req *commentdemo.DeleteCommentRequest) (resp *commentdemo.DeleteCommentResponse, err error) {
	// TODO: Your code here...
	resp = new(commentdemo.DeleteCommentResponse)

	if req.UserId < 0 || req.CommentId < 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewDeleteCommentService(ctx).DeleteComment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetVideoComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) GetVideoComment(ctx context.Context, req *commentdemo.GetVideoCommentRequest) (resp *commentdemo.GetVideoCommentResponse, err error) {
	// TODO: Your code here...
	resp = new(commentdemo.GetVideoCommentResponse)

	if req.UserId < 0 || req.VideoId < 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	comment_list, err := service.NewGetVideoCommetService(ctx).GetVideoComment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.CommentList = comment_list
	return resp, nil
}
