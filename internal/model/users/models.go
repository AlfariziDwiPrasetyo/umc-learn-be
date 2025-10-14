package users

import (
	"time"
)

type (
	RegisterUser struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Major    string `json:"major" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	LoginUser struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
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
