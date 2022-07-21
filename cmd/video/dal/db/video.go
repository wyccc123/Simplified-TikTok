package db

import (
	"context"
	"time"

	"github.com/cloudwego/simplified-tiktok/cmd/user/dal/db"
	constants "github.com/cloudwego/simplified-tiktok/pkg/constants"
	"github.com/cloudwego/simplified-tiktok/pkg/errno"
	"gorm.io/gorm"
)

// Video Gorm Data Structures
type Video struct {
	gorm.Model
	AuthorID      int    `gorm:"index:idx_authorid;not null"`
	PlayUrl       string `gorm:"type:varchar(255);not null"`
	CoverUrl      string `gorm:"type:varchar(255)"`
	FavoriteCount int    `gorm:"default:0"`
	CommentCount  int    `gorm:"default:0"`
	Title         string `gorm:"type:varchar(50);not null"`
}

func (Video) TableName() string {
	return constants.VideoTableName
}

// PublishVideo publish a new video
func PublishVideo(ctx context.Context, video *Video) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(video).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func DeleteVideo(ctx context.Context, videoID int64, userID int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Where("id = ? and user_id = ? ", videoID, userID).Delete(&Video{}).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func GetVideoList(ctx context.Context, limit int64, latest_time int64) ([]*Video, error) {
	video_list := make([]*Video, 0)

	conn := DB.WithContext(ctx)
	if err := conn.Limit(int(limit)).Order("update_time desc").Find(&video_list, "update_time < ?", time.UnixMilli(latest_time)).Error; err != nil {
		return nil, err
	}
	return video_list, nil
}

func GetUserPublished(ctx context.Context, userID int64) ([]*Video, error) {
	video_list := make([]*Video, 0)

	err := DB.WithContext(ctx).Model(&Video{}).Where(&Video{AuthorID: int(userID)}).Find(&video_list).Error
	if err != nil {
		return nil, err
	}

	return video_list, nil
}

func FavoriteVideo(ctx context.Context, userID int64, videoID int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		//1. 新增点赞数据
		user := new(db.User)
		if err := tx.WithContext(ctx).First(user, userID).Error; err != nil {
			return err
		}

		video := new(Video)
		if err := tx.WithContext(ctx).First(video, videoID).Error; err != nil {
			return err
		}

		if err := tx.WithContext(ctx).Model(&user).Association("FavoriteVideoIds").Append(videoID); err != nil {
			return err
		}
		//2.改变 video 表中的 favorite count
		res := tx.Model(video).Update("favorite_count", gorm.Expr("favorite_count + ?", 1))
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

func DisFavorite(ctx context.Context, userID int64, videoID int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		//1. 删除点赞数据
		user := new(db.User)
		if err := tx.WithContext(ctx).First(user, userID).Error; err != nil {
			return err
		}

		video := new(Video)
		if err := DB.WithContext(ctx).Model(&user).Association("FavoriteVideoIds").Find(&video, videoID); err != nil {
			return err
		}

		err := tx.Unscoped().WithContext(ctx).Model(&user).Association("FavoriteVideoIds").Delete(videoID)
		if err != nil {
			return err
		}

		//2.改变 video 表中的 favorite count
		res := tx.Model(video).Update("favorite_count", gorm.Expr("favorite_count - ?", 1))
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

func FavoriteList(ctx context.Context, userID int64) ([]*Video, error) {
	user := new(db.User)
	if err := DB.WithContext(ctx).First(user, userID).Error; err != nil {
		return nil, err
	}

	videoID_list := make([]int64, 0)
	if err := DB.WithContext(ctx).Model(&user).Association("FavoriteVideoIds").Find(&videoID_list); err != nil {
		return nil, err
	}
	videos := make([]*Video, 0)
	for _, vid := range videoID_list {
		video := new(Video)
		if err := DB.WithContext(ctx).First(video, vid).Error; err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}
	return videos, nil
}
