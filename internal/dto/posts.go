package dto

import (
	"time"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/posts"
)

type PostResponse struct {
	ID        int64             `json:"id"`
	Title     string            `json:"title"`
	Body      string            `json:"body"`
	Image     string            `json:"image,omitempty"`
	Slug      string            `json:"slug"`
	User      UserResponse      `json:"user"`
	Comments  []CommentResponse `json:"comments,omitempty"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

func ToPostResponse(p posts.Post) PostResponse {
	post := PostResponse{
		ID:        p.ID,
		Title:     p.Title,
		Body:      p.Body,
		Image:     p.Image,
		Slug:      p.Slug,
		User:      ToUserResponse(p.User),
		Comments:  ToCommentResponses(p.Comments),
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}

	return post
}

func ToPostResponses(list []posts.Post) []PostResponse {
	res := make([]PostResponse, len(list))
	for i, p := range list {
		res[i] = ToPostResponse(p)
	}

	return res
}
