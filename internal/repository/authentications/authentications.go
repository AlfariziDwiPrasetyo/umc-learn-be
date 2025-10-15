package authentications

import (
	"context"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/authentications"
)

func (r *Repository) StoreToken(ctx context.Context, model authentications.Authentications) error {
	err := r.Db.WithContext(ctx).Where("user_id = ?", model.UserID).Delete(&authentications.Authentications{}).Error
	if err != nil {
		return err
	}

	return r.Db.WithContext(ctx).Create(&model).Error
}

func (r *Repository) UpdateToken(ctx context.Context, userID int64, newToken string) error {
	return r.Db.WithContext(ctx).
		Model(&authentications.Authentications{}).
		Where("user_id = ?", userID).
		Update("refresh_token", newToken).Error
}

func (r *Repository) GetTokenByRefreshToken(ctx context.Context, refreshToken string) (*authentications.Authentications, error) {
	var auth authentications.Authentications

	if err := r.Db.WithContext(ctx).Where("refresh_token = ?", refreshToken).First(&auth).Error; err != nil {
		return nil, err
	}

	return &auth, nil
}

func (r *Repository) RevokeToken(ctx context.Context, refreshToken string) error {
	return r.Db.WithContext(ctx).
		Model(&authentications.Authentications{}).
		Where("refresh_token = ?", refreshToken).
		Update("revoked", true).Error
}
