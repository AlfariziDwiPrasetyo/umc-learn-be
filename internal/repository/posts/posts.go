package posts

import (
	"context"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/posts"
	"gorm.io/gorm"
)

func (r *Repository) CreatePost(ctx context.Context, post posts.Post) error {
	return r.Db.WithContext(ctx).Create(&post).Error
}

func (r *Repository) GetPosts(ctx context.Context, limit int) ([]posts.Post, error) {

	var posts []posts.Post
	if err := r.Db.WithContext(ctx).Limit(limit).Preload("User").Preload("Comments").
		Preload("Comments.User").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *Repository) GetPostById(ctx context.Context, id int64) (*posts.Post, error) {
	var post posts.Post
	err := r.Db.WithContext(ctx).
		Preload("User").
		Preload("Comments").
		Preload("Comments.User").
		Where("id = ?", id).First(&post).
		Error

	if err != nil {
		return nil, err
	}

	return &post, err
}

func (r *Repository) DeletePost(ctx context.Context, id int64) error {
	result := r.Db.WithContext(ctx).Delete(&posts.Post{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *Repository) UpdatePost(ctx context.Context, postID int64, updates map[string]any) error {
	if len(updates) == 0 {
		return nil
	}

	return r.Db.WithContext(ctx).Model(&posts.Post{}).Where("id = ?", postID).Updates(updates).Error
}
