package pack

import (
	"github.com/cloudwego/simplified-tiktok/cmd/user/dal/db"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/userdemo"
)

// User pack user info
func User(u *db.User) *userdemo.User {
	if u == nil {
		return nil
	}

	follow_count := int64(u.FollowCount)
	fans_count := int64(u.FansCount)
	isFollow := false

	return &userdemo.User{
		UserId:      int64(u.ID),
		UserName:    u.UserName,
		FollowCount: follow_count,
		FansCount:   fans_count,
		IsFollow:    isFollow,
	}
}

// Users pack list of user info
func Users(us []*db.User) []*userdemo.User {
	users := make([]*userdemo.User, 0)
	for _, u := range us {
		if user2 := User(u); user2 != nil {
			users = append(users, user2)
		}
	}
	return users
}
