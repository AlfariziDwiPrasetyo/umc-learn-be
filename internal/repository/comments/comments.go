package comments

import (
	"context"
	"errors"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/comments"
	"gorm.io/gorm"
)

func (r *Repository) CreateComment(ctx context.Context, comment comments.Comments) error {
	return r.Db.WithContext(ctx).Create(&comment).Error
}

func (r *Repository) GetCommentById(ctx context.Context, commentID int64) (*comments.Comments, error) {
	var comment comments.Comments
	err := r.Db.WithContext(ctx).First(&comment, commentID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("comment not found")
		}

		return nil, err
	}

	return &comment, nil

}

func (r *Repository) GetAllCommentsByPostId(ctx context.Context, postID int64) ([]comments.Comments, error) {
	var comments []comments.Comments

	err := r.Db.WithContext(ctx).Where("post_id = ?", postID).Preload("User").Order("created_at DESC").Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *Repository) DeleteComment(ctx context.Context, commentID int64) error {
	result := r.Db.WithContext(ctx).Delete(&comments.Comments{}, commentID)

	if result.RowsAffected == 0 {
		return errors.New("comment not found")
	}

	return result.Error
}

func (r *Repository) UpdateComment(ctx context.Context, commentID int64, updates map[string]any) error {
	if len(updates) == 0 {
		return nil
	}

	result := r.Db.WithContext(ctx).
		Model(&comments.Comments{}).
		Where("id = ?", commentID).
		Updates(updates)

	if result.RowsAffected == 0 {
		return errors.New("comment not found")
	}

	return result.Error
}
