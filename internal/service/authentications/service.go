package authentications

import (
	"context"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/configs"
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/authentications"
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/users"
)

type (
	authenticationsRepo interface {
		StoreToken(ctx context.Context, model authentications.Authentications) error
		UpdateToken(ctx context.Context, userID int64, newToken string) error
		GetTokenByRefreshToken(ctx context.Context, refreshToken string) (*authentications.Authentications, error)
		RevokeToken(ctx context.Context, refreshToken string) error
	}

	userRepo interface {
		CreateUser(ctx context.Context, model users.User) error
		GetUserById(ctx context.Context, id int64) (*users.User, error)
		GetUserByEmail(ctx context.Context, email string) (*users.User, error)
	}
)

type Service struct {
	Cfg            *configs.Config
	AuthRepository authenticationsRepo
	UserRepository userRepo
}

func NewService(cfg *configs.Config, userRepo userRepo, authRepo authenticationsRepo) *Service {
	return &Service{
		Cfg:            cfg,
		AuthRepository: authRepo,
		UserRepository: userRepo,
	}
}
