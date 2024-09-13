package dto

import (
	"time"
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
	MediaKeys *[]string `json:"media_keys,omitempty"`
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
