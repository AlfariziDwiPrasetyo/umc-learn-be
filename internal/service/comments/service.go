package comments

import (
	"context"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/configs"
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/comments"
	"github.com/cloudinary/cloudinary-go/v2"
)

type CommentsRepository interface {
	CreateComment(ctx context.Context, comment comments.Comments) error
	GetCommentById(ctx context.Context, commentID int64) (*comments.Comments, error)
	GetAllCommentsByPostId(ctx context.Context, postID int64) ([]comments.Comments, error)
	DeleteComment(ctx context.Context, commentID int64) error
	UpdateComment(ctx context.Context, commentID int64, updates map[string]any) error
}

type Service struct {
	CommentRepo CommentsRepository
	cfg         *configs.Config
	cld         *cloudinary.Cloudinary
}

func NewService(cfg *configs.Config, commentRepo CommentsRepository, cld *cloudinary.Cloudinary) *Service {
	return &Service{
		CommentRepo: commentRepo,
		cfg:         cfg,
		cld:         cld,
	}
}
