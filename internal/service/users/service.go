package users

import (
	"context"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/configs"
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/users"
)

type userRepo interface {
	CreateUser(ctx context.Context, model users.User) error
	GetUserById(ctx context.Context, id int64) (*users.User, error)
	GetUserByEmail(ctx context.Context, email string) (*users.User, error)
	DeleteUser(ctx context.Context, id int64) error
}

type Service struct {
	UserRepository userRepo
	Cfg            *configs.Config
}

func NewService(cfg *configs.Config, userRepository userRepo) *Service {
	return &Service{
		Cfg:            cfg,
		UserRepository: userRepository,
	}
}
