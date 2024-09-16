package service

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/bwmarrin/snowflake"
	"yabro.io/social-api/internal/db"
	"yabro.io/social-api/internal/db/postdb"
	"yabro.io/social-api/internal/logger"
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

type PostData struct {
	Post       postdb.Post
	Metrics    *postdb.PostPublicMetrics
	References []postdb.PostReference
	Tags       []postdb.PostTag
	Media      []postdb.PostMedia
}

func (s *PostService) GetPostByID(ctx context.Context, id int64, includeMetrics bool, includeReferences bool) (*PostData, error) {
	post, err := s.postDB.GetPost(ctx, id)
	if err != nil {
		return nil, err
	}

	var metrics *postdb.PostPublicMetrics
	if includeMetrics {
		metrics, err = s.postDB.GetPublicMetrics(ctx, id)
		if err != nil {
			return nil, err
		}
	}

	var references []postdb.PostReference
	if includeReferences {
		references, err = s.postDB.GetReferencesForPost(ctx, id)
		if err != nil {
			return nil, err
		}
	}

	return &PostData{
		Post:       *post,
		Metrics:    metrics,
		References: references,
	}, nil
}

type CreatePostParams struct {
	UserID        int64
	Content       string
	ReplyToPostID *int64
	QuotePostID   *int64
	Media         []*multipart.FileHeader
}

func (s *PostService) CreatePost(ctx context.Context, p CreatePostParams) error {
	id := s.snowflakeNode.Generate().Int64()

	var conversationID *int64
	var references []postdb.PostReference

	if p.ReplyToPostID != nil {
		replyToPost, err := s.postDB.GetPost(ctx, *p.ReplyToPostID)
		if err != nil {
			return fmt.Errorf("failed to get reply-to post: %w", err)
		}

		if replyToPost.ConversationID != nil {
			conversationID = replyToPost.ConversationID
		} else {
			conversationID = &replyToPost.ID
		}
		references = append(references, postdb.PostReference{
			SourcePostID:     id,
			ReferencedPostID: *p.ReplyToPostID,
			ReferenceType:    postdb.PostReferenceTypeReplyTo,
		})
	}

	if p.QuotePostID != nil {
		_, err := s.postDB.GetPost(ctx, *p.QuotePostID)
		if err != nil {
			return fmt.Errorf("failed to get quoted post: %w", err)
		}

		references = append(references, postdb.PostReference{
			SourcePostID:     id,
			ReferencedPostID: *p.QuotePostID,
			ReferenceType:    postdb.PostReferenceTypeQuote,
		})
	}

	tx, err := s.postDB.BeginTx(ctx)
	if err != nil {
		return err
	}

	createParams := postdb.CreatePostParams{
		ID:             id,
		Content:        p.Content,
		AuthorID:       p.UserID,
		ConversationID: conversationID,
	}

	// Create Post
	err = s.postDB.CreatePost(ctx, createParams, tx)
	if err != nil {
		return fmt.Errorf("failed to create post: %w", err)
	}

	// Create Metrics
	err = s.postDB.CreatePublicMetrics(ctx, postdb.PostPublicMetrics{PostID: id}, tx)
	if err != nil {
		return err
	}

	// Create References
	err = s.postDB.CreateReferences(ctx, references, tx)
	if err != nil {
		return err
	}

	if p.ReplyToPostID != nil {
		err = s.postDB.IncReplies(ctx, tx, *p.ReplyToPostID)
		if err != nil {
			return err
		}
	}
	if p.QuotePostID != nil {
		err = s.postDB.IncReposts(ctx, tx, *p.QuotePostID)
		if err != nil {
			return err
		}
	}

	// Create Media
	if len(p.Media) > 0 {
		media, err := s.uploadMedia(id, p.UserID, p.Media)
		if err != nil {
			return fmt.Errorf("failed to upload media: %w", err)
		}

		err = s.postDB.CreateMedia(ctx, media, tx)
		if err != nil {
			return fmt.Errorf("failed to add media to post: %w", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		logger.Err(ctx, err).Msg("CreatePost tx commit failed")
		return err
	}

	return nil
}

func (s *PostService) DeletePost(ctx context.Context, id int64, userID int64) error {
	err := s.postDB.DeletePost(ctx, id, userID, nil)
	if err != nil {
		return fmt.Errorf("failed to delete post: %w", err)
	}

	return nil
}

func (s *PostService) ListPosts(ctx context.Context, p postdb.ListPostParams) ([]PostData, *int64, error) {
	posts, nextCursor, err := s.postDB.ListPosts(ctx, p)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list posts: %w", err)
	}

	postDatas := make([]PostData, len(posts))
	postIDs := make([]int64, len(posts))
	for i, post := range posts {
		postIDs[i] = post.ID
		postDatas[i] = PostData{Post: post}
	}

	metrics, err := s.postDB.GetPublicMetricsForPosts(ctx, postIDs)
	if err != nil {
		return nil, nil, err
	}

	references, err := s.postDB.GetReferencesForPosts(ctx, postIDs)
	if err != nil {
		return nil, nil, err
	}

	for i, post := range postDatas {
		if m, ok := metrics[post.Post.ID]; ok {
			postDatas[i].Metrics = &m
		}
		if ref, ok := references[post.Post.ID]; ok {
			postDatas[i].References = ref
		}
	}

	return postDatas, nextCursor, nil
}
