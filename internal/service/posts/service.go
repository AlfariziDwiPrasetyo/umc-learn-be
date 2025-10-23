package posts

import (
	"context"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/configs"
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/posts"
	"github.com/cloudinary/cloudinary-go/v2"
)

type postRepository interface {
	CreatePost(ctx context.Context, post posts.Post) error
	GetPosts(ctx context.Context, limit int) ([]posts.Post, error)
	GetPostById(ctx context.Context, id int64) (*posts.Post, error)
	DeletePost(ctx context.Context, id int64) error
	UpdatePost(ctx context.Context, postID int64, updates map[string]any) error
}

type Service struct {
	postRepo postRepository
	cfg      *configs.Config
	cld      *cloudinary.Cloudinary
}

func NewService(cfg *configs.Config, postRepo postRepository, cld *cloudinary.Cloudinary) *Service {
	return &Service{
		postRepo: postRepo,
		cfg:      cfg,
		cld:      cld,
	}
}
