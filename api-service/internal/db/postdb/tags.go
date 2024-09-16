package postdb

import (
	"context"

	"github.com/jmoiron/sqlx"
	"yabro.io/social-api/internal/db/entity"
)

type CreatePostTagParams struct {
	PostID     int64   `db:"post_id"`
	EntityType string  `db:"entity_type"`
	StartIndex *int    `db:"start_index"`
	EndIndex   *int    `db:"end_index"`
	Tag        *string `db:"tag"`
}

type PostTag struct {
	ID int64 `db:"id"`
	CreatePostTagParams
}

func (pt PostTag) GetReferenceID() int64 { return pt.PostID }

func (pdb *PostDB) CreateTags(ctx context.Context, tags []CreatePostTagParams, tx *sqlx.Tx) error {
	exec := pdb.GetExecer(tx)
	return entity.CreateEntities[CreatePostTagParams](ctx, exec, "post_tags", tags)
}

func (pdb *PostDB) GetTagsForPosts(ctx context.Context, postIDs []int64) (map[int64][]PostTag, error) {
	return entity.GetEntitiesForIDs[PostTag](ctx, pdb.db, postIDs, "post_tags", "post_id", nil)
}
