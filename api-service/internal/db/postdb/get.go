package postdb

import (
	"database/sql"
	"fmt"
)

func (pdb *PostDB) GetPostData(id int64, includeMetrics bool) (*PostData, error) {
	// Query for the post
	postQuery := `SELECT * FROM posts WHERE id = $1 AND deleted_at IS NULL`
	var post Post
	err := pdb.db.Get(&post, postQuery, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("post not found")
		}
		return nil, fmt.Errorf("failed to get post: %w", err)
	}

	// Query for metrics if requested
	var metrics *PostPublicMetrics
	if includeMetrics {
		metricsQuery := `SELECT * FROM post_public_metrics WHERE post_id = $1`
		metrics = &PostPublicMetrics{}
		err = pdb.db.Get(metrics, metricsQuery, id)
		if err != nil && err != sql.ErrNoRows {
			return nil, fmt.Errorf("failed to get post metrics: %w", err)
		}
	}

	// Query for post references
	referencesQuery := `SELECT * FROM post_references WHERE source_post_id = $1`
	var references []PostReference
	err = pdb.db.Select(&references, referencesQuery, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to get post references: %w", err)
	}

	return &PostData{
		Post:       post,
		Metrics:    metrics,
		References: &references,
	}, nil
}
