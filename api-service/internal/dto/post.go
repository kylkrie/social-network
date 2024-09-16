package dto

import (
	"time"

	"yabro.io/social-api/internal/db/postdb"
	"yabro.io/social-api/internal/service"
	"yabro.io/social-api/internal/util"
)

type Post struct {
	ID             string             `json:"id"`
	Content        string             `json:"content"`
	AuthorID       string             `json:"author_id"`
	ConversationID *string            `json:"conversation_id,omitempty"`
	CreatedAt      time.Time          `json:"created_at"`
	IsDeleted      *bool              `json:"is_deleted,omitempty"`
	PublicMetrics  *PostPublicMetrics `json:"public_metrics,omitempty"`
	Edits          []PostEdit         `json:"edits,omitempty"`
	Tags           []PostTag          `json:"tags,omitempty"`
	References     []PostReference    `json:"references,omitempty"`
	Attachments    *PostAttachments   `json:"attachments,omitempty"`
}

type PostAttachments struct {
	MediaKeys []string `json:"media_keys,omitempty"`
}

type PostEdit struct {
	Content  string    `json:"content"`
	EditedAt time.Time `json:"edited_at"`
}

type PostTag struct {
	EntityType string  `json:"entity_type"`
	StartIndex *int    `json:"start_index,omitempty"`
	EndIndex   *int    `json:"end_index,omitempty"`
	Tag        *string `json:"tag,omitempty"`
}

type PostPublicMetrics struct {
	Reposts int `json:"reposts"`
	Replies int `json:"replies"`
	Likes   int `json:"likes"`
	Views   int `json:"views"`
}

type PostReference struct {
	ReferencedPostID string `json:"referenced_post_id"`
	ReferenceType    string `json:"reference_type"`
}

type Media struct {
	MediaKey string `json:"media_key"`
	Type     string `json:"type"`
	URL      string `json:"url"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
}

func ToPublicPost(p service.PostData, includes service.IncludeData) Post {
	if p.Post.DeletedAt != nil {
		deleted := true
		return Post{
			ID:             util.Int64ToString(p.Post.ID),
			Content:        "",
			AuthorID:       util.Int64ToString(p.Post.AuthorID),
			ConversationID: util.NullableInt64ToString(p.Post.ConversationID),
			CreatedAt:      p.Post.CreatedAt,
			IsDeleted:      &deleted,
		}
	}

	metrics := p.Metrics
	if metrics == nil {
		if includeMetrics, ok := includes.Metrics[p.Post.ID]; ok {
			metrics = &includeMetrics
		}
	}

	media := p.Media
	if media == nil {
		media = includes.Media[p.Post.ID]
	}

	return Post{
		ID:             util.Int64ToString(p.Post.ID),
		Content:        p.Post.Content,
		AuthorID:       util.Int64ToString(p.Post.AuthorID),
		ConversationID: util.NullableInt64ToString(p.Post.ConversationID),
		CreatedAt:      p.Post.CreatedAt,
		PublicMetrics:  toPublicPostMetrics(metrics),
		References:     toPublicPostReferences(p.References),
		Tags:           toPublicPostTags(p.Tags),
		Attachments:    toAttachments(media),
	}
}

func toPublicPostMetrics(metrics *postdb.PostPublicMetrics) *PostPublicMetrics {
	if metrics == nil {
		return nil
	}
	return &PostPublicMetrics{
		Reposts: metrics.Reposts,
		Replies: metrics.Replies,
		Likes:   metrics.Likes,
		Views:   metrics.Views,
	}
}

func toPublicPostReferences(references []postdb.PostReference) []PostReference {
	if len(references) == 0 {
		return nil
	}
	publicReferences := make([]PostReference, len(references))

	for i, ref := range references {
		publicReferences[i] = PostReference{
			ReferencedPostID: util.Int64ToString(ref.ReferencedPostID),
			ReferenceType:    ref.ReferenceType,
		}
	}
	return publicReferences
}

func toPublicPostTags(tags []postdb.PostTag) []PostTag {
	if len(tags) == 0 {
		return nil
	}
	publicTags := make([]PostTag, len(tags))
	for i, tag := range tags {
		publicTags[i] = PostTag{
			EntityType: tag.EntityType,
			StartIndex: tag.StartIndex,
			EndIndex:   tag.EndIndex,
			Tag:        tag.Tag,
		}
	}
	return publicTags
}

func toAttachments(media []postdb.PostMedia) *PostAttachments {
	if len(media) == 0 {
		return nil
	}

	mediaKeys := make([]string, len(media))
	for i, m := range media {
		mediaKeys[i] = util.Int64ToString(m.MediaKey)
	}

	return &PostAttachments{
		MediaKeys: mediaKeys,
	}
}

func ToPublicMedia(media postdb.PostMedia) Media {
	return Media{
		MediaKey: util.Int64ToString(media.MediaKey),
		Type:     media.Type,
		URL:      media.URL,

		Width:  media.Width,
		Height: media.Height,
	}
}
