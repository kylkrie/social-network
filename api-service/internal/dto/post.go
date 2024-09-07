package dto

import (
	"time"
)

type Post struct {
	ID             int64              `json:"id"`
	Content        string             `json:"content"`
	AuthorID       int64              `json:"author_id"`
	ConversationID *int64             `json:"conversation_id,omitempty"`
	CreatedAt      time.Time          `json:"created_at"`
	IsDeleted      *bool              `json:"is_deleted,omitempty"`
	PublicMetrics  *PostPublicMetrics `json:"public_metrics,omitempty"`
	Edits          []PostEdit         `json:"edits,omitempty"`
	Tags           []PostTag          `json:"tags,omitempty"`
	References     []PostReference    `json:"references,omitempty"`
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
	ReferencedPostID int64  `json:"referenced_post_id"`
	ReferenceType    string `json:"reference_type"`
}
