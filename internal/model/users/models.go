package users

import (
	"time"
)

type (
	UpdateUserRequest struct {
		Username *string `json:"username"`
		Email    *string `json:"email"`
		Major    *string `json:"major"`
		Image    *string `json:"image"`
		Password *string `json:"password"`
	}
)

type User struct {
	ID        int64  `gorm:"primaryKey;autoIncrement"`
	Username  string `gorm:"not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	Major     string
	Image     string
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
