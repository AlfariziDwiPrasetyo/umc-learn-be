package dto

import (
	"time"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/comments"
)

type CommentResponse struct {
	ID        int64        `json:"id"`
	PostID    int64        `json:"post_id"`
	Text      string       `json:"text"`
	Image     string       `json:"image,omitempty"`
	User      UserResponse `json:"user"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

func ToCommentResponse(c comments.Comments) CommentResponse {
	return CommentResponse{
		ID:        c.ID,
		PostID:    c.PostID,
		Text:      c.Text,
		Image:     c.Image,
		User:      ToUserResponse(c.User),
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func ToCommentResponses(list []comments.Comments) []CommentResponse {
	res := make([]CommentResponse, len(list))
	for i, c := range list {
		res[i] = ToCommentResponse(c)
	}
	return res
}
