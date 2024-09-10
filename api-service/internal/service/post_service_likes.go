package service

import (
	"fmt"

	"yabro.io/social-api/internal/dto"
	"yabro.io/social-api/internal/util"
)

func (s *PostService) LikePost(postID, userID int64) error {
	err := s.postDB.LikePost(postID, userID)
	if err != nil {
		return fmt.Errorf("failed to like post: %w", err)
	}
	return nil
}

func (s *PostService) UnlikePost(postID, userID int64) error {
	err := s.postDB.UnlikePost(postID, userID)
	if err != nil {
		return fmt.Errorf("failed to unlike post: %w", err)
	}
	return nil
}

func (s *PostService) ListUserLikes(userID int64, limit int, cursor *int64) ([]dto.Post, *string, error) {
	posts, nextCursor, err := s.postDB.ListUserLikes(userID, limit, cursor)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list user likes: %w", err)
	}

	publicPosts := make([]dto.Post, len(posts))
	for i, post := range posts {
		publicPosts[i] = *toPublicPost(post)
	}

	return publicPosts, util.NullableInt64ToString(nextCursor), nil
}
