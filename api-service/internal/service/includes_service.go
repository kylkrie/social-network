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

func (s *IncludeService) GetIncludesForPosts(posts []dto.Post, userID int64) (*dto.IncludeData, error) {
	// Collect unique IDs for original post set
	authorIDs := make(map[string]struct{})
	allPostIDs := make(map[string]struct{})
	otherPostIDs := make(map[string]struct{})

	for _, post := range posts {
		authorIDs[post.AuthorID] = struct{}{}
		allPostIDs[post.ID] = struct{}{} // Add original post IDs
		if post.ConversationID != nil {
			otherPostIDs[*post.ConversationID] = struct{}{}
			allPostIDs[*post.ConversationID] = struct{}{}
		}
		for _, ref := range post.References {
			otherPostIDs[ref.ReferencedPostID] = struct{}{}
			allPostIDs[ref.ReferencedPostID] = struct{}{}
		}
	}

	// Fetch include Posts
	uniqueOtherPostIDs := make([]int64, 0, len(otherPostIDs))
	for id := range otherPostIDs {
		uniqueOtherPostIDs = append(uniqueOtherPostIDs, util.StringToInt64MustParse(id))
	}
	includePosts, err := s.postDB.GetMany(uniqueOtherPostIDs)
	if err != nil {
		return nil, err
	}
	dtoPosts := make([]dto.Post, len(includePosts))
	for i, post := range includePosts {
		dtoPost := *toPublicPost(postdb.PostData{Post: post.Post, Metrics: &post.PublicMetrics})
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
	dtoUsers := make([]dto.User, len(includeUsers))
	for i, user := range includeUsers {
		dtoUsers[i] = toPublicUser(user, nil)
	}

	// Get user interactions for all unique post IDs
	allUniquePostIDs := make([]int64, 0, len(allPostIDs))
	for id := range allPostIDs {
		allUniquePostIDs = append(allUniquePostIDs, util.StringToInt64MustParse(id))
	}
	interactions, err := s.postDB.GetUserPostInteractions(allUniquePostIDs, userID)
	if err != nil {
		return nil, err
	}

	// Convert interactions to dtos
	dtoInteractions := make([]dto.UserPostInteraction, len(interactions))
	for i, interaction := range interactions {
		dtoInteractions[i] = *toPublicUserPostInteractions(&interaction)
	}

	// Fetch media
	mediaMap, err := s.postDB.GetMediaForPosts(allUniquePostIDs)
	if err != nil {
		return nil, err
	}

	var dtoMedia []dto.Media
	mediaKeyToMedia := make(map[string]dto.Media)
	postIDToMediaList := make(map[string][]string)

	for postID, mediaList := range mediaMap {
		postIDStr := util.Int64ToString(postID)
		for _, m := range mediaList {
			dtoM := toPublicMedia(&m)
			dtoMedia = append(dtoMedia, *dtoM)
			mediaKeyToMedia[dtoM.MediaKey] = *dtoM

			// Add media key to the post's list of media keys
			postIDToMediaList[postIDStr] = append(postIDToMediaList[postIDStr], dtoM.MediaKey)
		}
	}

	// Add attachments to dtoPosts
	for i, post := range dtoPosts {
		if mediaKeys, ok := postIDToMediaList[post.ID]; ok && len(mediaKeys) > 0 {
			dtoPosts[i].Attachments = &dto.PostAttachments{
				MediaKeys: &mediaKeys,
			}
		}
	}

	// Add attachments to original posts
	for i, post := range posts {
		if mediaKeys, ok := postIDToMediaList[post.ID]; ok && len(mediaKeys) > 0 {
			posts[i].Attachments = &dto.PostAttachments{
				MediaKeys: &mediaKeys,
			}
		}
	}

	return &dto.IncludeData{
		Users:            &dtoUsers,
		Posts:            &dtoPosts,
		UserInteractions: &dtoInteractions,
		Media:            &dtoMedia,
	}, nil
}

// Helper function to get includes for a single post
func (s *IncludeService) GetIncludesForPost(post *dto.Post, userID int64) (*dto.IncludeData, error) {
	includes, err := s.GetIncludesForPosts([]dto.Post{*post}, userID)
	if err != nil {
		return nil, err
	}
	return includes, nil
}
