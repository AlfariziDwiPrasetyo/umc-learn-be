package posts

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/posts"
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/utils"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

func (s *Service) GetPosts(ctx context.Context, limit int) ([]posts.Post, error) {
	if limit <= 0 {
		limit = 15
	}

	posts, err := s.postRepo.GetPosts(ctx, limit)

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *Service) CreatePost(ctx context.Context, userID int64, req posts.PostRequest) error {
	now := time.Now()
	id := time.Now().UnixNano()
	slug := fmt.Sprintf("%s-%d", slug.Make(req.Title), id)

	var imageUrl string
	if req.Image != nil {
		url, err := utils.UploadToCloudinary(ctx, req.Image, s.cld, s.cfg)
		if err != nil {
			return errors.New("failed to upload image")
		}

		imageUrl = url
	}

	post := posts.Post{
		ID:        id,
		Title:     req.Title,
		Body:      req.Body,
		Image:     imageUrl,
		UserID:    userID,
		Slug:      slug,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return s.postRepo.CreatePost(ctx, post)
}

func (s *Service) DeletePost(ctx context.Context, postID int64) error {
	post, err := s.postRepo.GetPostById(ctx, postID)
	if err != nil {
		return err
	}

	err = utils.DeleteFromCloudinary(ctx, s.cld, post.Image)
	if err != nil {
		return fmt.Errorf("failed to delete image : %s", err)
	}

	if post == nil {
		return errors.New("post not found")
	}

	return s.postRepo.DeletePost(ctx, postID)
}

func (s *Service) GetPostById(ctx context.Context, postID int64) (*posts.Post, error) {
	post, err := s.postRepo.GetPostById(ctx, postID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, err
	}

	return post, nil
}

func (s *Service) UpdatePost(ctx context.Context, postID int64, req posts.PostUpdateRequest) error {
	post, err := s.postRepo.GetPostById(ctx, postID)
	if err != nil {
		return err
	}

	if post == nil {
		return errors.New("post not found")
	}

	updates := make(map[string]any)

	if req.Title != nil {
		updates["title"] = *req.Title
	}
	if req.Body != nil {
		updates["body"] = *req.Body
	}

	if req.Image != nil {
		err := utils.DeleteFromCloudinary(ctx, s.cld, post.Image)

		if err != nil {
			return fmt.Errorf("failed to delete image: %w", err)
		}

		imageUrl, err := utils.UploadToCloudinary(ctx, req.Image, s.cld, s.cfg)

		updates["image"] = imageUrl
	}

	updates["updated_at"] = time.Now()

	return s.postRepo.UpdatePost(ctx, postID, updates)
}
