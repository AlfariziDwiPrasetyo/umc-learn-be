package dto

import "github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/users"

type UserResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email,omitempty"`
	Major    string `json:"major,omitempty"`
	Image    string `json:"image,omitempty"`
}

func ToUserResponse(u users.User) UserResponse {
	return UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Major:    u.Major,
		Image:    u.Image,
	}
}

func ToUserResponses(users []users.User) []UserResponse {
	res := make([]UserResponse, len(users))
	for i, u := range users {
		res[i] = ToUserResponse(u)
	}
	return res
}
