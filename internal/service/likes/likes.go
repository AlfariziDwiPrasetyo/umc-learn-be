package likes

import (
	"context"
	"fmt"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/likes"
	"github.com/rs/xid"
)

func (s *Service) CreateLike(ctx context.Context, postID int64, userID int64) error {
	liked, err := s.LikesRepo.IsPostLikedByUser(ctx, userID, postID)
	if err != nil {
		return fmt.Errorf("failed to get post liked by user : %w", err)
	}

	if liked {
		return fmt.Errorf("post already liked by user")
	}

	likeID := xid.New().String()

	like := likes.Like{
		ID:     likeID,
		PostID: postID,
		UserID: userID,
	}

	return s.LikesRepo.CreateLike(ctx, like)
}

func (s *Service) DeleteLike(ctx context.Context, postID int64, userID int64) error {
	liked, err := s.LikesRepo.IsPostLikedByUser(ctx, userID, postID)
	if err != nil {
		return fmt.Errorf("failed to get post liked by user : %w", err)
	}

	if !liked {
		return fmt.Errorf("liked not found")
	}

	return s.LikesRepo.DeleteLike(ctx, postID, userID)
}

func (s *Service) GetLikesByPostID(ctx context.Context, postID int64) ([]likes.Like, error) {
	likes, err := s.LikesRepo.GetLikesByPostID(ctx, postID)

	if err != nil {
		return nil, fmt.Errorf("failed to get likes : %w", err)
	}

	return likes, nil
}

func (s *Service) CountLikesByPostID(ctx context.Context, postID int64) (int64, error) {
	count, err := s.LikesRepo.CountLikesByPostID(ctx, postID)
	if err != nil {
		return 0, fmt.Errorf("failed to get count : %w", err)
	}

	return count, nil

}
