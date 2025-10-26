package likes

import (
	"context"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/configs"
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/likes"
)

type LikesRepository interface {
	CreateLike(ctx context.Context, like likes.Like) error
	DeleteLike(ctx context.Context, postID int64, userID int64) error
	IsPostLikedByUser(ctx context.Context, userID int64, postID int64) (bool, error)
	CountLikesByPostID(ctx context.Context, postID int64) (int64, error)
	GetLikesByPostID(ctx context.Context, postID int64) ([]likes.Like, error)
}

type Service struct {
	Cfg       *configs.Config
	LikesRepo LikesRepository
}

func NewService(cfg *configs.Config, likesRepo LikesRepository) *Service {
	return &Service{
		Cfg:       cfg,
		LikesRepo: likesRepo,
	}
}
