package comments

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/comments"
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/utils"
)

func (s *Service) CreateComment(ctx context.Context, userID int64, postID int64, req comments.CommentRequest) error {
	now := time.Now()
	commentID := time.Now().UnixNano()

	var imageUrl string

	if req.Image != nil {
		url, err := utils.UploadToCloudinary(ctx, req.Image, s.cld, s.cfg)
		if err != nil {
			return fmt.Errorf("failed to upload image : %s", err)
		}

		imageUrl = url

	}

	model := comments.Comments{
		ID:        commentID,
		UserID:    userID,
		PostID:    postID,
		Text:      req.Text,
		Image:     imageUrl,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return s.CommentRepo.CreateComment(ctx, model)

}

func (s *Service) GetAllCommentsByPostID(ctx context.Context, postID int64) ([]comments.Comments, error) {
	comments, err := s.CommentRepo.GetAllCommentsByPostId(ctx, postID)

	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (s *Service) GetCommentByID(ctx context.Context, commentsID int64) (*comments.Comments, error) {
	comment, err := s.CommentRepo.GetCommentById(ctx, commentsID)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *Service) UpdateComment(ctx context.Context, commentID int64, req comments.CommentRequest) error {

	comment, err := s.GetCommentByID(ctx, commentID)
	if err != nil {
		return err
	}

	if comment == nil {
		return errors.New("comment not found")
	}

	updates := make(map[string]any)

	if req.Text != "" {
		updates["text"] = req.Text
	}

	if req.Image != nil {
		err = utils.DeleteFromCloudinary(ctx, s.cld, comment.Image)
		if err != nil {
			return fmt.Errorf("failed to delete image : %s", err)
		}

		url, err := utils.UploadToCloudinary(ctx, req.Image, s.cld, s.cfg)
		if err != nil {
			return fmt.Errorf("failed to upload image : %s", err)
		}

		updates["image"] = url
	}

	updates["updated_at"] = time.Now()

	return s.CommentRepo.UpdateComment(ctx, commentID, updates)
}

func (s *Service) DeleteComment(ctx context.Context, commentID int64) error {
	comment, err := s.GetCommentByID(ctx, commentID)
	if err != nil {
		return err
	}

	if comment == nil {
		return errors.New("comment not found")
	}

	err = utils.DeleteFromCloudinary(ctx, s.cld, comment.Image)
	if err != nil {
		return fmt.Errorf("failed to delete image : %s", err)
	}

	return s.CommentRepo.DeleteComment(ctx, commentID)

}
