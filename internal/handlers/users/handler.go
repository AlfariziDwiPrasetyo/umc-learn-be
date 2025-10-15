package users

import (
	"context"

	"github.com/gin-gonic/gin"
)

type userService interface {
	DeleteUser(ctx context.Context, id int64) error
}

type Handler struct {
	*gin.Engine
	UserSvc userService
}

func NewHandler(api *gin.Engine, userSvc userService) *Handler {
	return &Handler{
		Engine:  api,
		UserSvc: userSvc,
	}
}

func (h *Handler) RegisterRoute() {
	// route := h.Group("users")
}
