package handlers

import (
	"context"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/simplified-tiktok/cmd/api/rpc"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/commentdemo"
	constants "github.com/cloudwego/simplified-tiktok/pkg/constants"
	"github.com/cloudwego/simplified-tiktok/pkg/errno"
	"github.com/gin-gonic/gin"
)

func GiveComment(c *gin.Context) {
	var commentVar CommentParam
	if err := c.ShouldBind(&commentVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))

	if userID < 0 || commentVar.VideoId < 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := rpc.GiveComment(context.Background(), &commentdemo.GiveCommentRequest{
		UserId:  userID,
		VideoId: commentVar.VideoId,
		Content: commentVar.Content,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

func DeleteComment(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))
	comment_id, err := strconv.ParseInt(c.Query("comment_id"), 64, 10)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if userID < 0 || comment_id < 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err1 := rpc.DeleteComment(context.Background(), &commentdemo.DeleteCommentRequest{
		UserId:    userID,
		CommentId: comment_id,
	})
	if err1 != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

func GetVideoComment(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))
	videoid, err := strconv.ParseInt(c.Query("video_id"), 64, 10)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if userID < 0 || videoid < 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	comment_list, err1 := rpc.GetVideoComment(context.Background(), &commentdemo.GetVideoCommentRequest{
		UserId:  userID,
		VideoId: videoid,
	})
	if err1 != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, comment_list)
}
