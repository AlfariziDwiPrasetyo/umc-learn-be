package authentications

import "time"

type (
	RefreshTokenRequest struct {
		RefreshToken string `json:"refresh_token"`
	}

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

type (
	Tokens struct {
		AccessToken  string
		RefreshToken string
	}
)

type Authentications struct {
	ID           int64  `gorm:"primaryKey;autoIncrement"`
	UserID       int64  `gorm:"not null;index"`
	RefreshToken string `gorm:"not null"`
	Revoked      bool   `gorm:"default:false"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
