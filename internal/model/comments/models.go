package comments

import (
	"mime/multipart"
	"time"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/users"
)

type (
	CommentRequest struct {
		Text  string                `form:"text"`
		Image *multipart.FileHeader `form:"image"`
	}
)

type Comments struct {
	ID        int64      `gorm:"primaryKey;autoIncrement"`
	UserID    int64      `gorm:"not null;index"`
	PostID    int64      `gorm:"not null;index"`
	Text      string     `gorm:"type:text;not null"`
	Image     string     `gorm:"type:text"`
	User      users.User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}
