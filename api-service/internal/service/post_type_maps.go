package service

import (
	"yabro.io/social-api/internal/db/postdb"
	"yabro.io/social-api/internal/dto"
)

func toPublicPost(post *postdb.Post, metrics *postdb.PostPublicMetrics) *dto.Post {
	if post.DeletedAt != nil {
		deleted := true
		return &dto.Post{
			ID:             post.ID,
			Content:        "",
			AuthorID:       post.AuthorID,
			ConversationID: post.ConversationID,
			CreatedAt:      post.CreatedAt,
			IsDeleted:      &deleted,
		}
	}

	publicPost := &dto.Post{
		ID:             post.ID,
		Content:        post.Content,
		AuthorID:       post.AuthorID,
		ConversationID: post.ConversationID,
		CreatedAt:      post.CreatedAt,
	}

	if metrics != nil {
		publicPost.PublicMetrics = &dto.PostPublicMetrics{
			Reposts: metrics.Reposts,
			Replies: metrics.Replies,
			Likes:   metrics.Likes,
			Views:   metrics.Views,
		}
	}

	return publicPost
}
