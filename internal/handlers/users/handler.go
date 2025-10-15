package users

import (
	"context"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/configs"
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/middleware"
	"github.com/gin-gonic/gin"
)

type userService interface {
	DeleteUser(ctx context.Context, id int64) error
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
}
