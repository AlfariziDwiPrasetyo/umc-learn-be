package authentications

import "time"

type (
	RefreshTokenRequest struct {
		RefreshToken string `json:"refresh_token"`
	}
)

type (
	Tokens struct {
		AccessToken  string
		RefreshToken string
	}
)

type AuthenticationModel struct {
	ID           int64  `gorm:"primaryKey;autoIncrement"`
	UserID       int64  `gorm:"not null;index"`
	RefreshToken string `gorm:"not null"`
	Revoked      bool   `gorm:"default:false"`
	ExpiredAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
