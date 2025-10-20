package comments

import (
	"time"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/users"
)

type (
	CommentRequest struct {
		Text  string `json:"text"`
		Image string `json:"image"`
	}
)

type Comments struct {
	ID        int64      `gorm:"primaryKey" binding:"required"`
	UserID    int64      `gorm:"user_id" binding:"required"`
	PostID    int64      `gorm:"post_id" binding:"required"`
	Text      string     `gorm:"text" binding:"required"`
	Image     string     `gorm:"image" binding:"required"`
	User      users.User `gorm:"foreignKey:UserID;references:ID"`
	CreatedAt time.Time  `gorm:"created_at" binding:"required"`
	UpdatedAt time.Time  `gorm:"updated_at" binding:"required"`
}
