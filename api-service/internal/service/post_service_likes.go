package service

import (
	"context"
	"fmt"
)

func (s *PostService) LikePost(ctx context.Context, postID, userID int64) error {
	err := s.postDB.LikePost(ctx, postID, userID)
	if err != nil {
		return fmt.Errorf("failed to like post: %w", err)
	}
	return nil
}

func (s *PostService) UnlikePost(ctx context.Context, postID, userID int64) error {
	err := s.postDB.UnlikePost(ctx, postID, userID)
	if err != nil {
		return fmt.Errorf("failed to unlike post: %w", err)
	}
	return nil
}

func (s *PostService) ListUserLikes(ctx context.Context, userID int64, limit int, cursor *int64) ([]PostData, *int64, error) {
	postsWithMetrics, nextCursor, err := s.postDB.ListUserLikes(ctx, userID, limit, cursor)
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
