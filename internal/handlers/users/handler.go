package users

import (
	"context"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/configs"
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/middleware"
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/users"
	"github.com/gin-gonic/gin"
)

type userService interface {
	DeleteUser(ctx context.Context, id int64) error
	UpdateUser(ctx context.Context, userID int64, req users.UpdateUserRequest) error
	GetUser(ctx context.Context, userID int64) (*users.User, error)
}

type Handler struct {
	*gin.Engine
	cfg     *configs.Config
	UserSvc userService
}

func NewHandler(api *gin.Engine, cfg *configs.Config, userSvc userService) *Handler {
	return &Handler{
		Engine:  api,
		UserSvc: userSvc,
		cfg:     cfg,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("users")

	route.Use(middleware.AuthMiddleware(h.cfg.Service.SecretKey))
	route.PATCH(":id", h.UpdateUser)
	route.GET(":id", h.GetUser)
	route.DELETE(":id", h.DeleteUser)
}
