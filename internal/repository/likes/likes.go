package likes

import (
	"context"
	"errors"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/likes"
)

func (r *Repository) CreateLike(ctx context.Context, like likes.Like) error {
	return r.Db.WithContext(ctx).Create(&like).Error
}

func (r *Repository) DeleteLike(ctx context.Context, postID int64, userID int64) error {
	result := r.Db.WithContext(ctx).Where("user_id = ? and post_id = ?", userID, postID).Delete(&likes.Like{})

	if result.RowsAffected == 0 {
		return errors.New("like not found")
	}

	return result.Error
}

func (r *Repository) IsPostLikedByUser(ctx context.Context, userID int64, postID int64) (bool, error) {
	var count int64
	err := r.Db.WithContext(ctx).Model(&likes.Like{}).Where("post_id = ? AND user_id = ?", postID, userID).Count(&count).Error
	return count > 0, err
}

func (r *Repository) CountLikesByPostID(ctx context.Context, postID int64) (int64, error) {
	var count int64
	err := r.Db.WithContext(ctx).Model(likes.Like{}).Where("post_id = ?", postID).Count(&count).Error

	return count, err
}

func (r *Repository) GetLikesByPostID(ctx context.Context, postID int64) ([]likes.Like, error) {
	var likeList []likes.Like
	err := r.Db.WithContext(ctx).Where("post_id", postID).Find(&likeList).Error

	return likeList, err
}
