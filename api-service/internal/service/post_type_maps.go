package service

import (
	"yabro.io/social-api/internal/db/postdb"
	"yabro.io/social-api/internal/dto"
	"yabro.io/social-api/internal/util"
)

func toPublicPost(p postdb.PostData) *dto.Post {
	if p.Post.DeletedAt != nil {
		deleted := true
		return &dto.Post{
			ID:             util.Int64ToString(p.Post.ID),
			Content:        "",
			AuthorID:       util.Int64ToString(p.Post.AuthorID),
			ConversationID: util.NullableInt64ToString(p.Post.ConversationID),
			CreatedAt:      p.Post.CreatedAt,
			IsDeleted:      &deleted,
		}
	}

	publicPost := &dto.Post{
		ID:             util.Int64ToString(p.Post.ID),
		Content:        p.Post.Content,
		AuthorID:       util.Int64ToString(p.Post.AuthorID),
		ConversationID: util.NullableInt64ToString(p.Post.ConversationID),
		CreatedAt:      p.Post.CreatedAt,
	}

	publicPost.PublicMetrics = toPublicPostMetrics(p.Metrics)
	publicPost.Edits = toPublicPostEdits(p.Edits)
	publicPost.References = toPublicPostReferences(p.References)
	publicPost.Tags = toPublicPostTags(p.Tags)

	return publicPost
}

func toPublicPostMetrics(metrics *postdb.PostPublicMetrics) *dto.PostPublicMetrics {
	if metrics == nil {
		return nil
	}
	return &dto.PostPublicMetrics{
		Reposts: metrics.Reposts,
		Replies: metrics.Replies,
		Likes:   metrics.Likes,
		Views:   metrics.Views,
	}
}

func toPublicPostEdits(edits *[]postdb.PostEdit) []dto.PostEdit {
	if edits == nil {
		return nil
	}
	publicEdits := make([]dto.PostEdit, len(*edits))
	for i, edit := range *edits {
		publicEdits[i] = dto.PostEdit{
			Content:  edit.Content,
			EditedAt: edit.EditedAt,
		}
	}
	return publicEdits
}

func toPublicPostReferences(references *[]postdb.PostReference) []dto.PostReference {
	if references == nil {
		return nil
	}
	publicReferences := make([]dto.PostReference, len(*references))
	for i, ref := range *references {
		publicReferences[i] = dto.PostReference{
			ReferencedPostID: util.Int64ToString(ref.ReferencedPostID),
			ReferenceType:    ref.ReferenceType,
		}
	}
	return publicReferences
}

func toPublicPostTags(tags *[]postdb.PostTag) []dto.PostTag {
	if tags == nil {
		return nil
	}
	publicTags := make([]dto.PostTag, len(*tags))
	for i, tag := range *tags {
		publicTags[i] = dto.PostTag{
			EntityType: tag.EntityType,
			StartIndex: tag.StartIndex,
			EndIndex:   tag.EndIndex,
			Tag:        tag.Tag,
		}
	}
	return publicTags
}

func toPublicMedia(media *postdb.PostMedia) *dto.Media {
	return &dto.Media{
		MediaKey: util.Int64ToString(media.MediaKey),
		Type:     media.Type,
		URL:      media.URL,
		Width:    media.Width,
		Height:   media.Height,
	}
}
