package postdb

import (
	"context"

	"github.com/jmoiron/sqlx"
	"yabro.io/social-api/internal/db/entity"
)

type PostReferenceType string

const (
	PostReferenceTypeQuote   = "quote"
	PostReferenceTypeReplyTo = "reply_to"
	PostReferenceTypeRepost  = "repost"
)

type PostReference struct {
	SourcePostID     int64  `db:"source_post_id"`
	ReferencedPostID int64  `db:"referenced_post_id"`
	ReferenceType    string `db:"reference_type"`
}

func (pt PostReference) GetReferenceID() int64 { return pt.SourcePostID }

func (pdb *PostDB) CreateReferences(ctx context.Context, references []PostReference, tx *sqlx.Tx) error {
	exec := pdb.GetExecer(tx)
	return entity.CreateEntities[PostReference](ctx, exec, "post_references", references)
}

func (pdb *PostDB) GetReferencesForPost(ctx context.Context, postID int64) ([]PostReference, error) {
	return entity.GetEntitiesForID[PostReference](ctx, pdb.db, postID, "post_references", "source_post_id")
}

func (pdb *PostDB) GetReferencesForPosts(ctx context.Context, postIDs []int64) (map[int64][]PostReference, error) {
	return entity.GetEntitiesForIDs[PostReference](ctx, pdb.db, postIDs, "post_references", "source_post_id", nil)
}
