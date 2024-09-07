package service

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
	"yabro.io/social-api/internal/db/postdb"
	"yabro.io/social-api/internal/dto"
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

func (s *PostService) GetPostByID(id int64, includeMetrics bool) (*dto.Post, error) {
	post, err := s.postDB.GetPostData(id, includeMetrics)
	if err != nil {
		return nil, fmt.Errorf("failed to get post: %w", err)
	}

	return toPublicPost(*post), nil
}

type CreatePostParams struct {
	UserID        int64
	Content       string
	ReplyToPostID *int64
	QuotePostID   *int64
}

func (s *PostService) CreatePost(p CreatePostParams) (*dto.Post, error) {
	id := s.snowflakeNode.Generate().Int64()

	var conversationID *int64
	var references []postdb.CreatePostReference

	if p.ReplyToPostID != nil {
		replyToPost, err := s.GetPostByID(*p.ReplyToPostID, false)
		if err != nil {
			return nil, fmt.Errorf("failed to get reply-to post: %w", err)
		}
		if replyToPost.ConversationID != nil {
			conversationID = replyToPost.ConversationID
		} else {
			conversationID = &replyToPost.ID
		}
		references = append(references, postdb.CreatePostReference{
			ID:              s.snowflakeNode.Generate().Int64(),
			ReferencePostID: *p.ReplyToPostID,
			ReferenceType:   postdb.PostReferenceTypeReplyTo,
		})
	}

	if p.QuotePostID != nil {
		_, err := s.GetPostByID(*p.QuotePostID, false)
		if err != nil {
			return nil, fmt.Errorf("failed to get quoted post: %w", err)
		}

		references = append(references, postdb.CreatePostReference{
			ID:              s.snowflakeNode.Generate().Int64(),
			ReferencePostID: *p.QuotePostID,
			ReferenceType:   postdb.PostReferenceTypeQuote,
		})
	}

	createParams := postdb.CreatePostParams{
		ID:             id,
		Content:        p.Content,
		AuthorID:       p.UserID,
		ConversationID: conversationID,
		References:     &references,
	}

	err := s.postDB.CreatePost(createParams)
	if err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	postData, err := s.postDB.GetPostData(id, false)
	if err != nil {
		return nil, err
	}

	return toPublicPost(*postData), nil
}

func (s *PostService) UpdatePost(id int64, userID int64, content string) (*dto.Post, error) {
	updateParams := postdb.UpdatePostParams{
		ID:      id,
		Content: content,
	}

	post, err := s.postDB.UpdatePost(updateParams)
	if err != nil {
		return nil, fmt.Errorf("failed to update post: %w", err)
	}

	return toPublicPost(postdb.PostData{Post: *post}), nil
}

func (s *PostService) DeletePost(id int64, userID int64) error {
	err := s.postDB.DeletePost(id, userID)
	if err != nil {
		return fmt.Errorf("failed to delete post: %w", err)
	}

	return nil
}

func (s *PostService) ListPosts(userID int64, limit int, cursor *int64) ([]dto.Post, *int64, error) {
	posts, nextCursor, err := s.postDB.ListPostDatas(userID, limit, cursor)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list posts: %w", err)
	}

	publicPosts := make([]dto.Post, len(posts))
	for i, post := range posts {
		publicPosts[i] = *toPublicPost(post)
	}

	return publicPosts, nextCursor, nil
}
