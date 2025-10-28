package likes

import (
	"context"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/configs"
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/middleware"
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/likes"
	"github.com/gin-gonic/gin"
)

type likesService interface {
	CreateLike(ctx context.Context, postID int64, userID int64) error
	DeleteLike(ctx context.Context, postID int64, userID int64) error
	GetLikesByPostID(ctx context.Context, postID int64) ([]likes.Like, error)
	CountLikesByPostID(ctx context.Context, postID int64) (int64, error)
}

type Handler struct {
	*gin.Engine
	LikesSvc likesService
	Cfg      *configs.Config
}

func NewHandler(api *gin.Engine, likesSvc likesService, cfg *configs.Config) *Handler {
	return &Handler{
		Engine:   api,
		LikesSvc: likesSvc,
		Cfg:      cfg,
	}
}

func (h *Handler) RegisterRoute() {
	r := h.Group("/posts")
	r.GET("/:id/likes", h.GetLikesByPostID)

	r.Use(middleware.AuthMiddleware(h.Cfg.Service.SecretKey))
	{
		r.POST("/:id/likes", h.CreateLike)
		r.DELETE("/:id/likes", h.DeleteLike)
	}
}
