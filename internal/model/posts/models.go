package posts

import (
	"mime/multipart"
	"time"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/comments"
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/users"
)

type (
	PostRequest struct {
		Title string                `form:"title" binding:"required"`
		Body  string                `form:"body" binding:"required"`
		Image *multipart.FileHeader `form:"image"`
	}

	PostUpdateRequest struct {
		Title *string               `form:"title"`
		Body  *string               `form:"body"`
		Image *multipart.FileHeader `form:"image"`
	}
)

type Post struct {
	ID        int64               `gorm:"primaryKey;autoIncrement"`
	Title     string              `gorm:"type:varchar(255);not null"`
	Body      string              `gorm:"type:text;not null"`
	UserID    int64               `gorm:"not null;index"`
	Image     string              `gorm:"type:text"`
	Slug      string              `gorm:"type:varchar(255);uniqueIndex"`
	User      users.User          `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Comments  []comments.Comments `gorm:"foreignKey:PostID;references:ID;constraint:OnDelete:CASCADE" json:"comments"`
	CreatedAt time.Time           `gorm:"autoCreateTime"`
	UpdatedAt time.Time           `gorm:"autoUpdateTime"`
}
