package handlers

import (
	"net/http"

	"github.com/cloudwego/simplified-tiktok/pkg/errno"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse pack response
func SendResponse(c *gin.Context, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
}

type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type GetUserParam struct {
	UserId int64 `json:"user_id,omitempty"` // 用户id
}

type FollowParam struct {
	FollowUserId int64 `json:"follow_user_id,omitempty"` // 对方用户id
	IsFollow     bool  `json:"is_follow,omitempty"`      // 1-关注，0-取消关注
}

type VideoParam struct {
	VideoUrl   string `json:"video_url,omitempty"`
	VideoTitle string `json:"video_title,omitempty"`
}

type FavoriteParam struct {
	VideoId    int64 `json:"video_id,omitempty"`
	IsFavorite bool  `json:"is_favorite,omitempty"`
}

type CommentParam struct {
	VideoId int64  `json:"video_id,omitempty"`
	Content string `json:"content,omitempty"`
}
