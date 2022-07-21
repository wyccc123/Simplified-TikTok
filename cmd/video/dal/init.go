package dal

import "github.com/cloudwego/simplified-tiktok/cmd/video/dal/db"

// Init init dal
func Init() {
	db.Init() // mysql init
}