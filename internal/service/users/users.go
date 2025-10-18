package users

import (
	"context"
	"errors"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/users"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *Service) DeleteUser(ctx context.Context, id int64) error {
	err := s.UserRepository.DeleteUser(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	return nil
}

func (s *Service) UpdateUser(ctx context.Context, userID int64, req users.UpdateUserRequest) error {
	updates := make(map[string]any)

	if req.Username != nil {
		updates["username"] = *req.Username
	}
	if req.Email != nil {
		updates["email"] = *req.Email
	}
	if req.Major != nil {
		updates["major"] = *req.Major
	}
	if req.Image != nil {
		updates["image"] = *req.Image
	}
	if req.Password != nil {
		pass, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		updates["password"] = pass
	}

	err := s.UserRepository.UpdateUser(ctx, userID, updates)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	return nil
}

func (s *Service) GetUser(ctx context.Context, userID int64) (*users.User, error) {
	user, err := s.UserRepository.GetUserById(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}
