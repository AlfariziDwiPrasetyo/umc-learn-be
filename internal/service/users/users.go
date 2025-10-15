package users

import (
	"context"
	"errors"

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
