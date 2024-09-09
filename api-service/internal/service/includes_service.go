package service

import (
	"yabro.io/social-api/internal/db/postdb"
	"yabro.io/social-api/internal/db/userdb"
	"yabro.io/social-api/internal/dto"
	"yabro.io/social-api/internal/util"
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

func (s *IncludeService) GetIncludesForPosts(posts []dto.Post) (*dto.IncludeData, error) {
	// Collect unique IDs for original post set
	authorIDs := make(map[string]struct{})
	postIDs := make(map[string]struct{})
	for _, post := range posts {
		authorIDs[post.AuthorID] = struct{}{}
		if post.ConversationID != nil {
			postIDs[*post.ConversationID] = struct{}{}
		}
		for _, ref := range post.References {
			postIDs[ref.ReferencedPostID] = struct{}{}
		}
	}

	// Fetch include Posts
	uniquePostIDs := make([]int64, 0, len(postIDs))
	for id := range postIDs {
		uniquePostIDs = append(uniquePostIDs, util.StringToInt64MustParse(id))
	}
	includePosts, err := s.postDB.GetMany(uniquePostIDs)
	if err != nil {
		return nil, err
	}
	dtoPosts := make([]dto.Post, len(includePosts))
	for i, post := range includePosts {
		dtoPost := *toPublicPost(postdb.PostData{Post: *post})
		dtoPosts[i] = dtoPost
		// add include post author IDs
		authorIDs[dtoPost.AuthorID] = struct{}{}
	}

	// Fetch includeUsers
	uniqueAuthorIDs := make([]int64, 0, len(authorIDs))
	for id := range authorIDs {
		uniqueAuthorIDs = append(uniqueAuthorIDs, util.StringToInt64MustParse(id))
	}
	includeUsers, err := s.userDB.GetMany(uniqueAuthorIDs)
	if err != nil {
		return nil, err
	}

	// Convert db users to dto users
	dtoUsers := make([]dto.User, len(includeUsers))
	for i, user := range includeUsers {
		dtoUsers[i] = toPublicUser(user, nil)
	}

	return &dto.IncludeData{
		Users: &dtoUsers,
		Posts: &dtoPosts,
	}, nil
}

// Helper function to get includes for a single post
func (s *IncludeService) GetIncludesForPost(post *dto.Post) (*dto.IncludeData, error) {
	includes, err := s.GetIncludesForPosts([]dto.Post{*post})
	if err != nil {
		return nil, err
	}
	return includes, nil
}
