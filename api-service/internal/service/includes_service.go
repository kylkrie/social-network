package service

import (
	"context"

	"yabro.io/social-api/internal/db/postdb"
	"yabro.io/social-api/internal/db/userdb"
)

type IncludeService struct {
	userDB *userdb.UserDB
	postDB *postdb.PostDB
}

func NewIncludeService(userDB *userdb.UserDB, postDB *postdb.PostDB) *IncludeService {
	return &IncludeService{
		userDB: userDB,
		postDB: postDB,
	}
}

type IncludeData struct {
	Posts        []postdb.Post
	Users        []userdb.User
	Metrics      map[int64]postdb.PostPublicMetrics
	Media        map[int64][]postdb.PostMedia
	Interactions map[int64]postdb.UserPostInteraction
}

func (s *IncludeService) GetIncludesForPosts(ctx context.Context, posts []PostData, userID *int64) (*IncludeData, error) {
	// Collect unique IDs for original post set
	authorIDs := make(map[int64]struct{})
	allPostIDs := make(map[int64]struct{})
	otherPostIDs := make(map[int64]struct{})

	for _, post := range posts {
		authorIDs[post.Post.AuthorID] = struct{}{}
		allPostIDs[post.Post.ID] = struct{}{} // Add original post IDs
		if post.Post.ConversationID != nil {
			otherPostIDs[*post.Post.ConversationID] = struct{}{}
			allPostIDs[*post.Post.ConversationID] = struct{}{}
		}
		for _, ref := range post.References {
			otherPostIDs[ref.ReferencedPostID] = struct{}{}
			allPostIDs[ref.ReferencedPostID] = struct{}{}
		}
	}

	// get unique IDs
	uniqueOtherPostIDs := make([]int64, 0, len(otherPostIDs))
	for id := range otherPostIDs {
		uniqueOtherPostIDs = append(uniqueOtherPostIDs, id)
	}
	allUniquePostIDs := make([]int64, 0, len(allPostIDs))
	for id := range allPostIDs {
		allUniquePostIDs = append(allUniquePostIDs, id)
	}
	uniqueAuthorIDs := make([]int64, 0, len(authorIDs))
	for id := range authorIDs {
		uniqueAuthorIDs = append(uniqueAuthorIDs, id)
	}

	// Fetch include Posts
	includePosts, err := s.postDB.GetPosts(ctx, uniqueOtherPostIDs)
	if err != nil {
		return nil, err
	}

	// Fetch metrics
	includePostMetrics, err := s.postDB.GetPublicMetricsForPosts(ctx, uniqueOtherPostIDs)
	if err != nil {
		return nil, err
	}

	// Fetch includeUsers
	includeUsers, err := s.userDB.GetMany(ctx, uniqueAuthorIDs)
	if err != nil {
		return nil, err
	}

	// Get user interactions for all unique post IDs
	var interactions map[int64]postdb.UserPostInteraction
	if userID != nil {
		interactions, err = s.postDB.GetUserPostInteractions(ctx, allUniquePostIDs, *userID)
		if err != nil {
			return nil, err
		}
	}

	// Fetch media
	mediaMap, err := s.postDB.GetMediaForPosts(ctx, allUniquePostIDs)
	if err != nil {
		return nil, err
	}

	return &IncludeData{
		Users:        includeUsers,
		Posts:        includePosts,
		Metrics:      includePostMetrics,
		Interactions: interactions,
		Media:        mediaMap,
	}, nil
}
