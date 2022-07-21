package db

import (
	"context"

	constants "github.com/cloudwego/simplified-tiktok/pkg/constants"
	"gorm.io/gorm"
)

// User Gorm Data Structures
type User struct {
	gorm.Model
	UserName         string  `gorm:"index:idx_username,unique;type:varchar(40);not null" json:"username"`
	Password         string  `gorm:"type:varchar(256);not null" json:"password"`
	FavoriteVideoIdS []int64 `json:"favorite_videoids"`
	FollowCount      int     `gorm:"default:0" json:"follow_count"`
	FansCount        int     `gorm:"default:0" json:"fans_count"`
}

func (User) TableName() string {
	return constants.UserTableName
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// GetUserByID multiple get list of user info
func GetUserByID(ctx context.Context, userID int64) (*User, error) {
	res := new(User)

	if err := DB.WithContext(ctx).First(&res, userID).Error; err != nil {
		return nil, err
	}
	return res, nil
}
