package likes

import "time"

type Like struct {
	ID        string    `gorm:"primaryKey"`
	PostID    int64     `gorm:"not null; index; uniqueIndex:idx_user_post"`
	UserID    int64     `gorm:"not null; index; uniqueIndex:idx_user_post"`
	CreatedAt time.Time `gorm:"created_at"`
}
