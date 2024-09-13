package service

import (
	"fmt"
	"mime/multipart"

	"github.com/bwmarrin/snowflake"
	"yabro.io/social-api/internal/db"
	"yabro.io/social-api/internal/db/postdb"
	"yabro.io/social-api/internal/dto"
	"yabro.io/social-api/internal/util"
)

type PostService struct {
	postDB        *postdb.PostDB
	snowflakeNode *snowflake.Node
	minioStorage  *db.MinioStorage
}

func NewPostService(postDB *postdb.PostDB, snowflakeNode *snowflake.Node, minioStorage *db.MinioStorage) (*PostService, error) {
	return &PostService{
		postDB:        postDB,
		snowflakeNode: snowflakeNode,
		minioStorage:  minioStorage,
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
	Media         []*multipart.FileHeader
}

func (s *PostService) CreatePost(p CreatePostParams) (*dto.Post, error) {
	id := s.snowflakeNode.Generate().Int64()

	var conversationID *string
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
		ConversationID: util.NullableStringToInt64MustParse(conversationID),
		References:     &references,
	}

	err := s.postDB.CreatePost(createParams)
	if err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	if len(p.Media) > 0 {
		media, err := s.uploadMedia(id, p.Media)
		if err != nil {
			return nil, fmt.Errorf("failed to upload media: %w", err)
		}

		err = s.postDB.AddMediaToPost(id, media)
		if err != nil {
			return nil, fmt.Errorf("failed to add media to post: %w", err)
		}
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

func (s *PostService) ListPosts(p postdb.ListPostParams) ([]dto.Post, *string, error) {
	posts, nextCursor, err := s.postDB.ListPostDatas(p)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list posts: %w", err)
	}

	publicPosts := make([]dto.Post, len(posts))
	for i, post := range posts {
		publicPosts[i] = *toPublicPost(post)
	}

	return publicPosts, util.NullableInt64ToString(nextCursor), nil
}
