package authentications

import (
	"context"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/authentications"
)

func (r *Repository) StoreToken(ctx context.Context, model authentications.AuthenticationModel) error {
	return r.Db.WithContext(ctx).Create(&model).Error
}

func (r *Repository) UpdateToken(ctx context.Context, userID int64, newToken string) error {
	return r.Db.WithContext(ctx).
		Model(&authentications.AuthenticationModel{}).
		Where("user_id = ?", userID).
		Update("refresh_token", newToken).Error
}

func (r *Repository) GetTokenByRefreshToken(ctx context.Context, refreshToken string) (*authentications.AuthenticationModel, error) {
	var auth authentications.AuthenticationModel

	if err := r.Db.WithContext(ctx).Where("refresh_token = ?", refreshToken).First(&auth).Error; err != nil {
		return nil, err
	}

	return &auth, nil
}

func (r *Repository) RevokeToken(ctx context.Context, refreshToken string) error {
	return r.Db.WithContext(ctx).
		Model(&authentications.AuthenticationModel{}).
		Where("refresh_token = ?", refreshToken).
		Update("revoked", true).Error
}
