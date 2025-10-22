package comments

import (
	"context"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/configs"
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/middleware"
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/comments"
	"github.com/gin-gonic/gin"
)

type CommentService interface {
	CreateComment(ctx context.Context, userID int64, postID int64, req comments.CommentRequest) error
	GetAllCommentsByPostID(ctx context.Context, postID int64) ([]comments.Comments, error)
	GetCommentByID(ctx context.Context, commentsID int64) (*comments.Comments, error)
	UpdateComment(ctx context.Context, commentID int64, req comments.CommentRequest) error
	DeleteComment(ctx context.Context, commentID int64) error
}

type Handler struct {
	*gin.Engine
	Cfg        *configs.Config
	CommentSvc CommentService
}

func NewHandler(api *gin.Engine, cfg *configs.Config, commentSvc CommentService) *Handler {
	return &Handler{
		Engine:     api,
		Cfg:        cfg,
		CommentSvc: commentSvc,
	}
}

func (h *Handler) RegisterRoute() {
	h.GET("/posts/:id/comments", h.GetAllCommentByPostID)

	r := h.Group("comments")
	r.Use(middleware.AuthMiddleware(h.Cfg.Service.SecretKey))
	{
		r.POST(":postID", h.CreateComment)
		r.PATCH(":commentID", h.UpdateComment)
		r.DELETE(":commentID", h.DeleteComment)
	}
}
