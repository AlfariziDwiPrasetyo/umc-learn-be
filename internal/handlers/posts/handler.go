package posts

import (
	"context"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/configs"
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/middleware"
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/posts"

	"github.com/gin-gonic/gin"
)

type postService interface {
	GetPosts(ctx context.Context, limit int) ([]posts.Post, error)
	CreatePost(ctx context.Context, userID int64, req posts.PostRequest) error
	DeletePost(ctx context.Context, postID int64) error
	GetPostById(ctx context.Context, postID int64) (*posts.Post, error)
	UpdatePost(ctx context.Context, postID int64, req posts.PostUpdateRequest) error
}

type Handler struct {
	*gin.Engine
	Cfg     *configs.Config
	PostSvc postService
}

func NewHandler(api *gin.Engine, cfg *configs.Config, postSvc postService) *Handler {
	return &Handler{
		Engine:  api,
		Cfg:     cfg,
		PostSvc: postSvc,
	}
}

func (h *Handler) RegisterRoute() {
	r := h.Group("posts")

	r.Use(middleware.AuthMiddleware(h.Cfg.Service.SecretKey))

	r.POST("", h.CreatePost)
	r.PATCH(":id", h.UpdatePost)
	r.DELETE(":id", h.DeletePost)
	r.GET(":id", h.GetPostById)
	r.GET("", h.GetAllPost)
}
