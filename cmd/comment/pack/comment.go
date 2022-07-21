package pack

import (
	"github.com/cloudwego/simplified-tiktok/cmd/comment/dal/db"
	"github.com/cloudwego/simplified-tiktok/kitex_gen/commentdemo"
)

func Comment(c *db.Comment) *commentdemo.Comment {
	if c == nil {
		return nil
	}

	return &commentdemo.Comment{
		CommentId:  int64(c.ID),
		UserId:     int64(c.UserID),
		Content:    c.Content,
		CreateDate: c.CreatedAt.String(),
	}
}

// Comments pack list of comment info
func Comments(cs []*db.Comment) []*commentdemo.Comment {
	comments := make([]*commentdemo.Comment, 0)
	for _, c := range cs {
		if comment := Comment(c); comment != nil {
			comments = append(comments, comment)
		}
	}
	return comments
}
