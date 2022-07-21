package pack

import (
	"github.com/cloudwego/simplified-tiktok/cmd/video/dal/db"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/videodemo"
)

// Video pack video info
func Video(v *db.Video) *videodemo.Video {
	if v == nil {
		return nil
	}

	return &videodemo.Video{
		VideoId:       int64(v.ID),
		UserId:        int64(v.AuthorID),
		VideoUrl:      v.PlayUrl,
		VideoCoverUrl: v.CoverUrl,
		FavoriteCount: int64(v.FavoriteCount),
		CommentCount:  int64(v.CommentCount),
		VideoTitle:    v.Title,
		IsFavorite:    false,
		CreateTime:    v.CreatedAt.Unix(),
	}
}

// Videos pack list of video info
func Videos(vs []*db.Video) []*videodemo.Video {
	videos := make([]*videodemo.Video, 0)
	for _, v := range vs {
		if n := Video(v); n != nil {
			videos = append(videos, n)
		}
	}
	return videos
}
