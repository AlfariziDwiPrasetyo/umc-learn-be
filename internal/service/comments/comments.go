package comments

import (
	"context"
	"time"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/comments"
)

func (s *Service) CreateComment(ctx context.Context, userID int64, postID int64, req comments.CommentRequest) error {
	now := time.Now()
	commentID := time.Now().UnixNano()

	model := comments.Comments{
		ID:        commentID,
		UserID:    userID,
		PostID:    postID,
		Text:      req.Text,
		Image:     req.Image,
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
	updates := make(map[string]any)

	if req.Image != "" {
		updates["image"] = req.Image
	}

	if req.Text != "" {
		updates["text"] = req.Text
	}

	err := s.CommentRepo.UpdateComment(ctx, commentID, updates)

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteComment(ctx context.Context, commentID int64) error {

	return s.CommentRepo.DeleteComment(ctx, commentID)

}
