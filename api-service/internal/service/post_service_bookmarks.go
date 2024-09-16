package service

import (
	"context"
	"fmt"
)

func (s *PostService) BookmarkPost(ctx context.Context, postID, userID int64) error {
	_, err := s.postDB.BookmarkPost(ctx, postID, userID, nil)
	if err != nil {
		return fmt.Errorf("failed to bookmark post: %w", err)
	}

	return nil
}

func (s *PostService) UnbookmarkPost(ctx context.Context, postID, userID int64) error {
	_, err := s.postDB.UnbookmarkPost(ctx, postID, userID, nil)
	if err != nil {
		return fmt.Errorf("failed to unbookmark post: %w", err)
	}
	return nil
}

func (s *PostService) ListUserBookmarks(ctx context.Context, userID int64, limit int, cursor *int64) ([]PostData, *int64, error) {
	postsWithMetrics, nextCursor, err := s.postDB.ListUserBookmarks(ctx, userID, limit, cursor)
	if err != nil {
		return nil, nil, err
	}

	postDatas := make([]PostData, len(postsWithMetrics))
	for i, post := range postsWithMetrics {
		postDatas[i] = PostData{
			Post:    post.Post,
			Metrics: &post.Metrics,
		}
	}

	return postDatas, nextCursor, nil
}
