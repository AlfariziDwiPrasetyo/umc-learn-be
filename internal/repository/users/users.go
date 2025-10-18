package users

import (
	"context"
	"errors"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/users"
	"gorm.io/gorm"
)

func (r *Repository) CreateUser(ctx context.Context, model users.User) error {
	result := r.Db.WithContext(ctx).Create(&model)

	return result.Error
}

func (r *Repository) GetUserById(ctx context.Context, id int64) (*users.User, error) {
	var user users.User
	err := r.Db.WithContext(ctx).First(&user, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*users.User, error) {
	var user users.User
	err := r.Db.WithContext(ctx).Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) DeleteUser(ctx context.Context, id int64) error {
	result := r.Db.WithContext(ctx).Delete(&users.User{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *Repository) UpdateUser(ctx context.Context, userID int64, updates map[string]any) error {
	if len(updates) == 0 {
		return nil
	}

	return r.Db.WithContext(ctx).
		Model(&users.User{}).
		Where("id = ?", userID).
		Updates(updates).Error
}
