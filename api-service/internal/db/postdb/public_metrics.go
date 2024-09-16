package postdb

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"yabro.io/social-api/internal/db/entity"
	"yabro.io/social-api/internal/util"
)

type PostPublicMetrics struct {
	PostID  int64 `db:"post_id"`
	Reposts int   `db:"reposts"`
	Replies int   `db:"replies"`
	Likes   int   `db:"likes"`
	Views   int   `db:"views"`
}

type UpdatePostPublicMetricsParams struct {
	Reposts *int `db:"reposts"`
	Replies *int `db:"replies"`
	Likes   *int `db:"likes"`
	Views   *int `db:"views"`
}

func (pdb *PostDB) CreatePublicMetrics(ctx context.Context, params PostPublicMetrics, tx *sqlx.Tx) error {
	exec := pdb.GetExecer(tx)
	return entity.CreateEntity(ctx, exec, "post_public_metrics", params)
}

func (pdb *PostDB) GetPublicMetrics(ctx context.Context, postID int64) (*PostPublicMetrics, error) {
	return entity.GetEntity[PostPublicMetrics](ctx, pdb.db, postID, "post_public_metrics", "post_id")
}

func (pdb *PostDB) GetPublicMetricsForPosts(ctx context.Context, postIDs []int64) (map[int64]PostPublicMetrics, error) {
	metrics, err := entity.GetEntities[PostPublicMetrics](ctx, pdb.db, postIDs, "post_public_metrics", "post_id")
	if err != nil {
		return nil, err
	}

	return util.ArrToMap(metrics, func(m PostPublicMetrics) int64 { return m.PostID }), nil
}

func (pdb *PostDB) IncLikes(ctx context.Context, tx *sqlx.Tx, postID int64) error {
	return pdb.increaseMetric(ctx, tx, "likes", postID)
}

func (pdb *PostDB) DecLikes(ctx context.Context, tx *sqlx.Tx, postID int64) error {
	return pdb.decreaseMetric(ctx, tx, "likes", postID)
}

func (pdb *PostDB) IncReplies(ctx context.Context, tx *sqlx.Tx, postID int64) error {
	return pdb.increaseMetric(ctx, tx, "replies", postID)
}

func (pdb *PostDB) DecReplies(ctx context.Context, tx *sqlx.Tx, postID int64) error {
	return pdb.decreaseMetric(ctx, tx, "replies", postID)
}

func (pdb *PostDB) IncReposts(ctx context.Context, tx *sqlx.Tx, postID int64) error {
	return pdb.increaseMetric(ctx, tx, "reposts", postID)
}

func (pdb *PostDB) DecReposts(ctx context.Context, tx *sqlx.Tx, postID int64) error {
	return pdb.decreaseMetric(ctx, tx, "reposts", postID)
}

func (pdb *PostDB) increaseMetric(ctx context.Context, tx *sqlx.Tx, columnName string, postID int64) error {
	exec := pdb.GetExecer(tx)

	updateQuery := fmt.Sprintf(`
            INSERT INTO post_public_metrics (post_id, %s)
            VALUES ($1, 1)
            ON CONFLICT (post_id)
            DO UPDATE SET %s = post_public_metrics.%s + 1
        `, columnName, columnName, columnName)
	_, err := exec.ExecContext(ctx, updateQuery, postID)
	if err != nil {
		return fmt.Errorf("failed to increase post metrics %s: %w", columnName, err)
	}

	return nil
}

func (pdb *PostDB) decreaseMetric(ctx context.Context, tx *sqlx.Tx, columnName string, postID int64) error {
	exec := pdb.GetExecer(tx)

	updateQuery := fmt.Sprintf(`
            UPDATE post_public_metrics
            SET %s = GREATEST(likes - 1, 0)
            WHERE post_id = $1
        `, columnName)
	_, err := exec.ExecContext(ctx, updateQuery, postID)
	if err != nil {
		return fmt.Errorf("failed to decrease post metrics %s: %w", columnName, err)
	}

	return nil
}
