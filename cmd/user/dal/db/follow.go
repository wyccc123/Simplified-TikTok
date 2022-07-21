package db

import (
	"context"

	"gorm.io/gorm"

	constants "github.com/cloudwego/simplified-tiktok/pkg/constants"
	"github.com/cloudwego/simplified-tiktok/pkg/errno"
)

type Follow struct {
	gorm.Model
	User         User `gorm:"foreignkey:UserID;"`
	UserID       int  `gorm:"index:idx_userid,unique;not null"`
	FollowUser   User `gorm:"foreignkey:FollowUserID;"`
	FollowUserID int  `gorm:"index:idx_userid,unique;index:idx_userid_to;not null"`
}

func (Follow) TableName() string {
	return constants.VideoTableName
}

func NewFollow(ctx context.Context, userID int64, followUserID int64) error {
	// 开启事务操作
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		//1. 新增关注数据
		err := tx.Create(&Follow{UserID: int(userID), FollowUserID: int(followUserID)}).Error
		if err != nil {
			return err
		}

		//2.改变 user 表中的 follow count
		res := tx.Model(new(User)).Where("ID = ?", userID).Update("follow_count", gorm.Expr("follow_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.ErrDatabase
		}

		return nil
	})
	return err
}

func DisFollow(ctx context.Context, userID int64, followUserID int64) error {
	// 开启事务操作
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		follow := new(Follow)
		if err := tx.Where("user_id = ? AND follow_user_id=?", userID, followUserID).First(&follow).Error; err != nil {
			return err
		}

		//1. 删除关注数据
		err := tx.Unscoped().Delete(&follow).Error
		if err != nil {
			return err
		}
		//2.改变 user 表中的 following count
		res := tx.Model(new(User)).Where("ID = ?", userID).Update("follow_count", gorm.Expr("follow_count - ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.ErrDatabase
		}

		return nil
	})
	return err
}

func GetFollowList(ctx context.Context, userID int64) ([]*User, error) {
	var follow_list []*Follow
	err := DB.WithContext(ctx).Where("user_id = ?", userID).Find(&follow_list).Error
	if err != nil {
		return nil, err
	}
	user_list := make([]*User, 0)
	for _, f := range follow_list {
		user, err := GetUserByID(ctx, int64(f.FollowUserID))
		if err != nil {
			return nil, err
		}
		user_list = append(user_list, user)
	}
	return user_list, nil
}
