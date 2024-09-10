package postdb

import (
	"fmt"
	"strings"
)

type PostWithMetrics struct {
	Post
	PublicMetrics PostPublicMetrics `db:"metrics"`
}

func (pdb *PostDB) GetMany(ids []int64) ([]*PostWithMetrics, error) {
	if len(ids) == 0 {
		return []*PostWithMetrics{}, nil
	}

	// Create a string of placeholders for the SQL query
	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = id
	}

	query := fmt.Sprintf(`
		SELECT 
			p.*,
			COALESCE(m.reposts, 0) AS "metrics.reposts",
			COALESCE(m.replies, 0) AS "metrics.replies",
			COALESCE(m.likes, 0) AS "metrics.likes",
			COALESCE(m.views, 0) AS "metrics.views"
		FROM posts p
		LEFT JOIN post_public_metrics m ON p.id = m.post_id
		WHERE p.id IN (%s) AND p.deleted_at IS NULL
	`, strings.Join(placeholders, ","))

	var postsWithMetrics []*PostWithMetrics
	err := pdb.db.Select(&postsWithMetrics, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get posts with metrics: %w", err)
	}

	return postsWithMetrics, nil
}
