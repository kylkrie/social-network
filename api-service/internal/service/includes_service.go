package service

import (
	"yabro.io/social-api/internal/db/postdb"
	"yabro.io/social-api/internal/db/userdb"
	"yabro.io/social-api/internal/dto"
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
	// Collect unique author IDs
	authorIDs := make(map[int64]struct{})
	for _, post := range posts {
		authorIDs[post.AuthorID] = struct{}{}
	}

	// Convert map to slice
	uniqueAuthorIDs := make([]int64, 0, len(authorIDs))
	for id := range authorIDs {
		uniqueAuthorIDs = append(uniqueAuthorIDs, id)
	}

	// Fetch users
	users, err := s.userDB.GetMany(uniqueAuthorIDs)
	if err != nil {
		return nil, err
	}

	// Convert db users to dto users
	dtoUsers := make([]dto.User, len(users))
	for i, user := range users {
		dtoUsers[i] = toPublicUser(user, nil)
	}

	return &dto.IncludeData{
		Users: &dtoUsers,
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
