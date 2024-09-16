package postdb

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

func (pdb *PostDB) BookmarkPost(ctx context.Context, postID, userID int64, tx *sqlx.Tx) (bool, error) {
	exec := pdb.GetExecer(tx)
	query := `
        INSERT INTO post_bookmarks (post_id, user_id)
        VALUES ($1, $2)
        ON CONFLICT (post_id, user_id) DO NOTHING
    `
	result, err := exec.ExecContext(ctx, query, postID, userID)
	if err != nil {
		return false, fmt.Errorf("failed to bookmark post: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func (pdb *PostDB) UnbookmarkPost(ctx context.Context, postID, userID int64, tx *sqlx.Tx) (bool, error) {
	exec := pdb.GetExecer(tx)
	query := `
        DELETE FROM post_bookmarks
        WHERE post_id = $1 AND user_id = $2
    `

	result, err := exec.ExecContext(ctx, query, postID, userID)
	if err != nil {
		return false, fmt.Errorf("failed to unbookmark post: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func (pdb *PostDB) ListUserBookmarks(ctx context.Context, userID int64, limit int, cursor *int64) ([]PostWithMetrics, *int64, error) {
	query := strings.Builder{}
	query.WriteString(`
        SELECT p.*, m.reposts, m.replies, m.likes, m.views
        FROM posts p
        JOIN post_bookmarks b ON p.id = b.post_id
        LEFT JOIN post_public_metrics m ON p.id = m.post_id
        WHERE b.user_id = $1 AND p.deleted_at IS NULL
    `)

	args := []interface{}{userID}

	if cursor != nil {
		query.WriteString(fmt.Sprintf(" AND p.id < $%d", len(args)+1))
		args = append(args, *cursor)
	}

	query.WriteString(" ORDER BY p.id DESC")
	query.WriteString(fmt.Sprintf(" LIMIT $%d", len(args)+1))

	args = append(args, limit+1)

	rows, err := pdb.db.QueryxContext(ctx, query.String(), args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list user bookmarks: %w", err)
	}
	defer rows.Close()

	var posts []PostWithMetrics
	for rows.Next() {
		var post Post
		var metrics PostPublicMetrics
		err := rows.Scan(
			&post.ID, &post.Content, &post.AuthorID, &post.ConversationID,
			&post.CreatedAt, &post.UpdatedAt, &post.DeletedAt,
			&metrics.Reposts, &metrics.Replies, &metrics.Likes, &metrics.Views,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to scan post: %w", err)
		}
		posts = append(posts, PostWithMetrics{Post: post, Metrics: metrics})
	}

	var nextCursor *int64
	if len(posts) > limit {
		nextCursor = &posts[limit-1].Post.ID
		posts = posts[:limit]
	}

	return posts, nextCursor, nil
}
