package comments

import "time"

type (
	CommentRequest struct {
		Text  string `json:"text"`
		Image string `json:"image"`
	}
)

type Comments struct {
	ID        int64     `gorm:"primaryKey" binding:"required"`
	UserID    int64     `gorm:"user_id" binding:"required"`
	PostID    int64     `gorm:"post_id" binding:"required"`
	Text      string    `gorm:"text" binding:"required"`
	Image     string    `gorm:"image" binding:"required"`
	CreatedAt time.Time `gorm:"created_at" binding:"required"`
	UpdatedAt time.Time `gorm:"updated_at" binding:"required"`
}
