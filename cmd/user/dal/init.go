package dal

import "github.com/cloudwego/simplified-tiktok/cmd/user/dal/db"

// Init init dal
func Init() {
	db.Init() // mysql init
}
