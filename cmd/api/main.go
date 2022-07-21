package api

import (
	"context"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/simplified-tiktok/cmd/api/handlers"
	"github.com/cloudwego/simplified-tiktok/cmd/api/rpc"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/userdemo"
	constants "github.com/cloudwego/simplified-tiktok/pkg/constants"
	"github.com/cloudwego/simplified-tiktok/pkg/tracer"
	"github.com/gin-gonic/gin"
)

func Init() {
	tracer.InitJaeger(constants.ApiServiceName)
	rpc.InitRPC()
}

func main() {
	Init()
	r := gin.New()
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVar handlers.UserParam
			if err := c.ShouldBind(&loginVar); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
				return "", jwt.ErrMissingLoginValues
			}

			return rpc.Login(context.Background(), &userdemo.RegisterRequest{UserName: loginVar.UserName, Password: loginVar.PassWord})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	douyin := r.Group("/douyin")

	user := douyin.Group("/user")
	user.POST("/login/", authMiddleware.LoginHandler)
	user.POST("/register/", handlers.Register)
	user.GET("/", handlers.GetUserById)

	video := douyin.Group("/feed")
	video.Use(authMiddleware.MiddlewareFunc())
	video.GET("/", handlers.GetVideoList)

	publish := douyin.Group("/publish")
	publish.Use(authMiddleware.MiddlewareFunc())
	publish.POST("/action/", handlers.Publish)
	publish.GET("/list/", handlers.GetUserPublished)

	favorite := douyin.Group("/favorite")
	favorite.Use(authMiddleware.MiddlewareFunc())
	favorite.POST("/action/", handlers.FavoriteVideo)
	favorite.GET("/list/", handlers.GetUserFavorite)

	comment := douyin.Group("/comment")
	comment.Use(authMiddleware.MiddlewareFunc())
	comment.POST("/action/", handlers.GetVideoComment)
	comment.GET("/list/", handlers.GetVideoComment)

	relation := douyin.Group("/relation")
	relation.Use(authMiddleware.MiddlewareFunc())
	relation.POST("/action/", handlers.Follow)
	relation.GET("/follow/list/", handlers.GetFollowList)

	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}
