package postdb

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func (pdb *PostDB) ListPosts(userID int64, limit int, cursor *int64) ([]Post, *int64, error) {
	query := `
		SELECT *
		FROM posts
		WHERE author_id = $1 AND deleted_at IS NULL
	`
	args := []interface{}{userID}

	if cursor != nil {
		query += ` AND id < $2
		ORDER BY id DESC
		LIMIT $3`
		args = append(args, *cursor, limit+1) // fetch one extra to determine if there are more results
	} else {
		query += `
		ORDER BY id DESC
		LIMIT $2`
		args = append(args, limit+1) // fetch one extra to determine if there are more results
	}

	rows, err := pdb.db.Queryx(query, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list posts: %w", err)
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.StructScan(&post); err != nil {
			return nil, nil, fmt.Errorf("failed to scan post: %w", err)
		}
		posts = append(posts, post)
	}

	var nextCursor *int64
	if len(posts) > limit {
		nextCursor = &posts[limit-1].ID
		posts = posts[:limit]
	}

	return posts, nextCursor, nil
}

func (pdb *PostDB) ListPostDatas(userID int64, limit int, cursor *int64) ([]PostData, *int64, error) {
	// First, fetch the posts
	posts, nextCursor, err := pdb.ListPosts(userID, limit, cursor)
	if err != nil {
		return nil, nil, err
	}

	if len(posts) == 0 {
		return []PostData{}, nextCursor, nil
	}

	// Collect post IDs
	postIDs := make([]int64, len(posts))
	for i, post := range posts {
		postIDs[i] = post.ID
	}

	// Fetch metrics
	metricsQuery := `SELECT * FROM post_public_metrics WHERE post_id IN (?)`
	query, args, err := sqlx.In(metricsQuery, postIDs)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create IN query for metrics: %w", err)
	}
	query = pdb.db.Rebind(query)

	var metrics []PostPublicMetrics
	err = pdb.db.Select(&metrics, query, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch metrics: %w", err)
	}

	// Fetch references
	referencesQuery := `SELECT * FROM post_references WHERE source_post_id IN (?)`
	query, args, err = sqlx.In(referencesQuery, postIDs)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create IN query for references: %w", err)
	}
	query = pdb.db.Rebind(query)

	var references []PostReference
	err = pdb.db.Select(&references, query, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch references: %w", err)
	}

	// Organize metrics and references by post ID
	metricsMap := make(map[int64]*PostPublicMetrics)
	for i := range metrics {
		metricsMap[metrics[i].PostID] = &metrics[i]
	}

	referencesMap := make(map[int64][]PostReference)
	for _, ref := range references {
		referencesMap[ref.SourcePostID] = append(referencesMap[ref.SourcePostID], ref)
	}

	// Combine everything
	result := make([]PostData, len(posts))
	for i, post := range posts {
		postData := PostData{
			Post:    post,
			Metrics: metricsMap[post.ID],
		}

		if len(referencesMap) > 0 {
			refs := referencesMap[post.ID]
			postData.References = &refs
		}

		result[i] = postData
	}

	return result, nextCursor, nil
}
