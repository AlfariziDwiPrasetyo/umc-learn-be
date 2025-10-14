package users

import (
	"context"
	"errors"
	"time"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/users"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) SignUp(ctx context.Context, req users.RegisterUser) error {
	user, err := s.UserRepository.GetUserByEmail(ctx, req.Email)

	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("email already exist")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	now := time.Now()

	model := users.User{
		Username:  req.Username,
		Email:     req.Email,
		Major:     req.Major,
		Password:  string(pass),
		CreatedAt: now,
		UpdatedAt: now,
	}

	return s.UserRepository.CreateUser(ctx, model)
}
