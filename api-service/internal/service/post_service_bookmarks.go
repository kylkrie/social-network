package service

import (
	"fmt"

	"yabro.io/social-api/internal/dto"
	"yabro.io/social-api/internal/util"
)

func (s *PostService) BookmarkPost(postID, userID int64) error {
	err := s.postDB.BookmarkPost(postID, userID)
	if err != nil {
		return fmt.Errorf("failed to bookmark post: %w", err)
	}

	return nil
}

func (s *PostService) UnbookmarkPost(postID, userID int64) error {
	err := s.postDB.UnbookmarkPost(postID, userID)
	if err != nil {
		return fmt.Errorf("failed to unbookmark post: %w", err)
	}
	return nil
}

func (s *PostService) ListUserBookmarks(userID int64, limit int, cursor *int64) ([]dto.Post, *string, error) {
	posts, nextCursor, err := s.postDB.ListUserBookmarks(userID, limit, cursor)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list user bookmarks: %w", err)
	}

	publicPosts := make([]dto.Post, len(posts))
	for i, post := range posts {
		publicPosts[i] = *toPublicPost(post)
	}

	return publicPosts, util.NullableInt64ToString(nextCursor), nil
}
