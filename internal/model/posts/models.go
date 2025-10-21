package posts

import (
	"time"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/comments"
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/users"
)

type (
	PostRequest struct {
		Title string `json:"title" binding:"required"`
		Body  string `json:"body" binding:"required"`
		Image string `json:"image"`
	}
	PostUpdateRequest struct {
		Title *string `json:"title"`
		Body  *string `json:"body"`
		Image *string `json:"image"`
	}
)

type Post struct {
	ID        int64               `gorm:"primaryKey" binding:"required"`
	Title     string              `gorm:"title" binding:"required"`
	Body      string              `gorm:"body" binding:"required"`
	UserID    int64               `gorm:"user_id" binding:"required"`
	Image     string              `gorm:"image"`
	Slug      string              `gorm:"slug" binding:"required"`
	User      users.User          `gorm:"foreignKey:UserID;references:ID"`
	Comments  []comments.Comments `gorm:"foreignKey:PostID;references:ID" json:"comments"`
	CreatedAt time.Time           `gorm:"created_at" binding:"required"`
	UpdatedAt time.Time           `gorm:"updated_at" binding:"required"`
}
