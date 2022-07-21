package dal

import "github.com/cloudwego/simplified-tiktok/cmd/comment/dal/db"

// Init init dal
func Init() {
	db.Init() // mysql init
}
