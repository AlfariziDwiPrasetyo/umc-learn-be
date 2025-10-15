package authentications

import (
	"context"
	"errors"
	"time"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/authentications"
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/users"
	"github.com/alfarizidwiprasetyo/be-umc-learn/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) StoreToken(ctx context.Context, refreshToken string, userID int64) error {

	now := time.Now()

	model := authentications.Authentications{
		UserID:       userID,
		RefreshToken: refreshToken,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	return s.AuthRepository.StoreToken(ctx, model)
}

func (s *Service) SignIn(ctx context.Context, req authentications.LoginUser) (*authentications.Tokens, error) {
	user, err := s.UserRepository.GetUserByEmail(ctx, req.Email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := jwt.CreateToken(user.ID, s.Cfg.Service.SecretKey)

	if err != nil {
		return nil, err
	}

	err = s.StoreToken(ctx, token.RefreshToken, user.ID)

	if err != nil {
		return nil, err
	}

	return token, nil

}

func (s *Service) SignUp(ctx context.Context, req authentications.RegisterUser) error {
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

func (s *Service) Refresh(ctx context.Context, req authentications.RefreshTokenRequest) (*authentications.Tokens, error) {
	claims, err := jwt.ValidateToken(string(req.RefreshToken), s.Cfg.Service.SecretKey)

	if err != nil {
		return nil, err
	}

	token, err := jwt.CreateToken(claims, s.Cfg.Service.SecretKey)

	if err != nil {
		return nil, err
	}

	return token, nil

}
