package postdb

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"yabro.io/social-api/internal/db/entity"
)

type CreatePostMediaParams struct {
	MediaKey int64  `db:"media_key"`
	PostID   int64  `db:"post_id"`
	UserID   int64  `db:"user_id"`
	Type     string `db:"type"`
	URL      string `db:"url"`
	Width    int    `db:"width"`
	Height   int    `db:"height"`
}

type PostMedia struct {
	CreatePostMediaParams
	CreatedAt time.Time `db:"created_at"`
}

func (pm PostMedia) GetReferenceID() int64 { return pm.PostID }

func (pdb *PostDB) CreateMedia(ctx context.Context, media []CreatePostMediaParams, tx *sqlx.Tx) error {
	exec := pdb.GetExecer(tx)
	return entity.CreateEntities(ctx, exec, "post_media", media)
}

func (pdb *PostDB) GetMediaForPosts(ctx context.Context, postIDs []int64) (map[int64][]PostMedia, error) {
	idColName := "media_key"
	return entity.GetEntitiesForIDs[PostMedia](ctx, pdb.db, postIDs, "post_media", "post_id", &idColName)
}
