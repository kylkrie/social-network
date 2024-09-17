package service

import (
	"context"
	"fmt"

	"yabro.io/social-api/internal/db/postdb"
)

type ListRepliesForPostParams struct {
	PostID int64
	Limit  int

	Cursor *int64
}

func (s *PostService) ListRepliesForPost(ctx context.Context, params ListRepliesForPostParams) ([]PostData, *int64, error) {
	replies, nextCursor, err := s.postDB.ListRepliesForPost(ctx, postdb.ListRepliesForPostParams{
		PostID: params.PostID,
		Limit:  params.Limit,
		Cursor: params.Cursor,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list replies for post: %w", err)
	}

	postDatas := make([]PostData, len(replies))
	postIDs := make([]int64, len(replies))
	for i, reply := range replies {
		postIDs[i] = reply.ID
		postDatas[i] = PostData{Post: reply}
	}

	// Fetch metrics for all replies
	metrics, err := s.postDB.GetPublicMetricsForPosts(ctx, postIDs)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get metrics for replies: %w", err)
	}

	// Fetch references for all replies
	references, err := s.postDB.GetReferencesForPosts(ctx, postIDs)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get references for replies: %w", err)
	}

	// Enrich PostData with fetched information
	for i, postData := range postDatas {

		if m, ok := metrics[postData.Post.ID]; ok {
			postDatas[i].Metrics = &m
		}
		if refs, ok := references[postData.Post.ID]; ok {
			postDatas[i].References = refs
		}

	}

	return postDatas, nextCursor, nil
}
