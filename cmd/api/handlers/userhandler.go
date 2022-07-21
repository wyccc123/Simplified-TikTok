package handlers

import (
	"context"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	constants "github.com/cloudwego/simplified-tiktok/pkg/constants"
	"github.com/cloudwego/simplified-tiktok/pkg/errno"

	"github.com/cloudwego/simplified-tiktok/kitex_gen/userdemo"

	"github.com/cloudwego/simplified-tiktok/cmd/api/rpc"

	"github.com/gin-gonic/gin"
)

// Register register user info
func Register(c *gin.Context) {
	var registerVar UserParam
	if err := c.ShouldBind(&registerVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	_, err := rpc.Register(context.Background(), &userdemo.RegisterRequest{
		UserName: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

func Login(c *gin.Context) {
	var loginVar UserParam
	if err := c.ShouldBind(&loginVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	_, err := rpc.Register(context.Background(), &userdemo.RegisterRequest{
		UserName: loginVar.UserName,
		Password: loginVar.PassWord,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

func GetUserById(c *gin.Context) {
	var getuserVar GetUserParam
	userid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	getuserVar.UserId = int64(userid)
	if getuserVar.UserId < 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	user, err := rpc.GetUserById(context.Background(), &userdemo.GetUserRequest{
		UserId: getuserVar.UserId,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, user)
}

func Follow(c *gin.Context) {
	var followVar FollowParam
	if err := c.ShouldBind(&followVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))

	if followVar.FollowUserId < 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	if userID < 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := rpc.Follow(context.Background(), &userdemo.FollowRequest{
		UserId:       userID,
		FollowUserId: followVar.FollowUserId,
		IsFollow:     followVar.IsFollow,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

func GetFollowList(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))

	if userID < 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	userlist, err := rpc.GetFollowList(context.Background(), &userdemo.GetFollowListRequest{
		UserId: userID,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, userlist)
}
