package postdb

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type PostDB struct {
	db *sqlx.DB
}

func NewPostDB(db *sqlx.DB) *PostDB {
	return &PostDB{db: db}
}

type PostData struct {
	Post       Post
	Edits      *[]PostEdit
	Metrics    *PostPublicMetrics
	References *[]PostReference
	Tags       *[]PostTag
	Media      *[]PostMedia
}

type Post struct {
	ID             int64      `db:"id"`
	Content        string     `db:"content"`
	AuthorID       int64      `db:"author_id"`
	ConversationID *int64     `db:"conversation_id"`
	CreatedAt      time.Time  `db:"created_at"`
	UpdatedAt      time.Time  `db:"updated_at"`
	DeletedAt      *time.Time `db:"deleted_at"`
}

type PostMedia struct {
	MediaKey  int64     `db:"media_key"`
	PostID    int64     `db:"post_id"`
	Type      string    `db:"type"`
	URL       string    `db:"url"`
	Width     int       `db:"width"`
	Height    int       `db:"height"`
	CreatedAt time.Time `db:"created_at"`
}

type PostEdit struct {
	ID       int64     `db:"id"`
	PostID   int64     `db:"post_id"`
	Content  string    `db:"content"`
	EditedAt time.Time `db:"edited_at"`
}

type PostTag struct {
	ID         int64   `db:"id"`
	PostID     int64   `db:"post_id"`
	EntityType string  `db:"entity_type"`
	StartIndex *int    `db:"start_index"`
	EndIndex   *int    `db:"end_index"`
	Tag        *string `db:"tag"`
}

type PostPublicMetrics struct {
	PostID  int64 `db:"post_id"`
	Reposts int   `db:"reposts"`
	Replies int   `db:"replies"`
	Likes   int   `db:"likes"`
	Views   int   `db:"views"`
}

type PostReference struct {
	ID               int64  `db:"id"`
	SourcePostID     int64  `db:"source_post_id"`
	ReferencedPostID int64  `db:"referenced_post_id"`
	ReferenceType    string `db:"reference_type"`
}

type PostReferenceType string

const (
	PostReferenceTypeQuote   = "quote"
	PostReferenceTypeReplyTo = "reply_to"
	PostReferenceTypeRepost  = "repost"
)
