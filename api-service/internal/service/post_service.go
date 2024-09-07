package service

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
	"yabro.io/social-api/internal/db/postdb"
)

type PostService struct {
	postDB        *postdb.PostDB
	snowflakeNode *snowflake.Node
}

func NewPostService(postDB *postdb.PostDB, snowflakeNode *snowflake.Node) (*PostService, error) {
	return &PostService{
		postDB:        postDB,
		snowflakeNode: snowflakeNode,
	}, nil
}

func (s *PostService) GetPostByID(id int64, includeMetrics bool) (*Post, error) {
	post, metrics, err := s.postDB.GetPost(id, includeMetrics)
	if err != nil {
		return nil, fmt.Errorf("failed to get post: %w", err)
	}

	return toPublicPost(post, metrics), nil
}

func (s *PostService) CreatePost(userID int64, content string, conversationID *int64) (*Post, error) {
	id := s.snowflakeNode.Generate().Int64()

	createParams := postdb.CreatePostParams{
		ID:             id,
		Content:        content,
		AuthorID:       userID,
		ConversationID: conversationID,
	}

	post, err := s.postDB.CreatePost(createParams)
	if err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	return toPublicPost(post, nil), nil
}

func (s *PostService) UpdatePost(id int64, userID int64, content string) (*Post, error) {
	updateParams := postdb.UpdatePostParams{
		ID:      id,
		Content: content,
	}

	post, err := s.postDB.UpdatePost(updateParams)
	if err != nil {
		return nil, fmt.Errorf("failed to update post: %w", err)
	}

	return toPublicPost(post, nil), nil
}

func (s *PostService) DeletePost(id int64, userID int64) error {
	err := s.postDB.DeletePost(id, userID)
	if err != nil {
		return fmt.Errorf("failed to delete post: %w", err)
	}

	return nil
}

func (s *PostService) ListPosts(userID int64, limit int, cursor *int64) ([]Post, *int64, error) {
	posts, nextCursor, err := s.postDB.ListPosts(userID, limit, cursor)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list posts: %w", err)
	}

	publicPosts := make([]Post, len(posts))
	for i, post := range posts {
		publicPosts[i] = *toPublicPost(&post, nil)
	}

	return publicPosts, nextCursor, nil
}
