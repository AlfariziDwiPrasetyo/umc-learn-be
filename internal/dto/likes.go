package dto

import (
	"time"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/likes"
)

type LikeResponse struct {
	ID        string    `json:"id"`
	PostID    int64     `json:"post_id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func ToLikeResponse(l likes.Like) LikeResponse {
	like := LikeResponse{
		ID:        l.ID,
		PostID:    l.PostID,
		UserID:    l.UserID,
		CreatedAt: l.CreatedAt,
	}

	return like
}

func ToLikeResponses(list []likes.Like) []LikeResponse {
	res := make([]LikeResponse, len(list))
	for i, p := range list {
		res[i] = ToLikeResponse(p)
	}

	return res
}
