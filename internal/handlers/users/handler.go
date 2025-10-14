package users

import (
	"context"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/users"
	"github.com/gin-gonic/gin"
)

type userService interface {
	SignUp(ctx context.Context, req users.RegisterUser) error
}

type Handler struct {
	*gin.Engine
	UserSvc userService
}

func NewHandler(api *gin.Engine, userService userService) *Handler {
	return &Handler{
		Engine:  api,
		UserSvc: userService,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("users")

	route.POST("/sign-up", h.SignUp)
}
