package handlers

import (
	"context"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/simplified-tiktok/cmd/api/rpc"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/videodemo"
	constants "github.com/cloudwego/simplified-tiktok/pkg/constants"
	"github.com/cloudwego/simplified-tiktok/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Publish(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))

	if userID < 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := rpc.Publish(context.Background(), &videodemo.PublishRequest{
		UserId: userID,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

func DeleteVideo(c *gin.Context) {
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

	if err := rpc.DeleteVideo(context.Background(), &videodemo.DeleteVideoRequest{
		UserId:  userID,
		VideoId: videoid,
	}); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

func GetVideoList(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))
	latest_time, err := strconv.ParseInt(c.Query("video_id"), 64, 10)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if userID < 0 || latest_time < 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	video_list, next_time, err1 := rpc.GetVideoList(context.Background(), &videodemo.GetVideoListRequst{
		UserId:     userID,
		LatestTime: latest_time,
	})
	if err1 != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, map[string]interface{}{"video_list": video_list, "next_time": next_time})
}

func GetUserPublished(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))

	if userID < 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	video_list, err := rpc.GetUserPublished(context.Background(), &videodemo.GetUserPublishedRequest{
		UserId: userID,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, video_list)
}

func FavoriteVideo(c *gin.Context) {
	var favoriteVar FavoriteParam
	if err := c.ShouldBind(&favoriteVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))

	if userID < 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := rpc.FavoriteVideo(context.Background(), &videodemo.FavoriteVideoRequest{
		UserId:     userID,
		VideoId:    favoriteVar.VideoId,
		IsFavorite: favoriteVar.IsFavorite,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

func GetUserFavorite(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))

	if userID < 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	video_list, err := rpc.GetUserFavorite(context.Background(), &videodemo.GetUserFavoriteRequest{
		UserId: userID,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, video_list)
}
