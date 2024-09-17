package postdb

import (
	"context"
	"fmt"
)

type ListRepliesForPostParams struct {
	PostID int64
	Limit  int
	Cursor *int64
}

func (pdb *PostDB) ListRepliesForPost(ctx context.Context, params ListRepliesForPostParams) ([]Post, *int64, error) {
	query := `
		SELECT p.*
		FROM posts p
		JOIN post_references pr ON p.id = pr.source_post_id
		WHERE pr.referenced_post_id = $1
		  AND pr.reference_type = 'reply_to'
	`
	args := []interface{}{params.PostID}

	if params.Cursor != nil {
		query += " AND p.id < $2"
		args = append(args, *params.Cursor)
	}

	// Always add the LIMIT clause
	query += fmt.Sprintf(" ORDER BY p.id DESC LIMIT %d", params.Limit+1)

	rows, err := pdb.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list replies for post: %w", err)
	}
	defer rows.Close()

	var replies []Post
	for rows.Next() {
		var reply Post
		if err := rows.StructScan(&reply); err != nil {
			return nil, nil, fmt.Errorf("failed to scan reply post: %w", err)
		}
		replies = append(replies, reply)
	}

	if err = rows.Err(); err != nil {
		return nil, nil, fmt.Errorf("error iterating reply posts: %w", err)
	}

	var nextCursor *int64
	if len(replies) > params.Limit {
		nextCursor = &replies[params.Limit-1].ID
		replies = replies[:params.Limit]
	}

	return replies, nextCursor, nil
}
