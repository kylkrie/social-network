package postdb

import (
	"context"
	"fmt"
	"strings"
)

func (pdb *PostDB) LikePost(ctx context.Context, postID, userID int64) error {
	tx, err := pdb.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	query := `
        INSERT INTO post_likes (post_id, user_id)
        VALUES ($1, $2)
        ON CONFLICT (post_id, user_id) DO NOTHING
    `
	result, err := tx.ExecContext(ctx, query, postID, userID)
	if err != nil {
		return fmt.Errorf("failed to like post: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected > 0 {
		// Only update metrics if a new like was actually inserted
		updateQuery := `
            INSERT INTO post_public_metrics (post_id, likes)
            VALUES ($1, 1)
            ON CONFLICT (post_id)
            DO UPDATE SET likes = post_public_metrics.likes + 1
        `
		_, err = tx.ExecContext(ctx, updateQuery, postID)
		if err != nil {
			return fmt.Errorf("failed to update post metrics: %w", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (pdb *PostDB) UnlikePost(ctx context.Context, postID, userID int64) error {
	tx, err := pdb.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	query := `
        DELETE FROM post_likes
        WHERE post_id = $1 AND user_id = $2
    `
	result, err := tx.ExecContext(ctx, query, postID, userID)
	if err != nil {
		return fmt.Errorf("failed to unlike post: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected > 0 {
		// Only update metrics if a like was actually removed
		updateQuery := `
            UPDATE post_public_metrics
            SET likes = GREATEST(likes - 1, 0)
            WHERE post_id = $1
        `
		_, err = tx.ExecContext(ctx, updateQuery, postID)
		if err != nil {
			return fmt.Errorf("failed to update post metrics: %w", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (pdb *PostDB) ListUserLikes(ctx context.Context, userID int64, limit int, cursor *int64) ([]PostWithMetrics, *int64, error) {
	query := strings.Builder{}
	query.WriteString(`
        SELECT p.*, m.reposts, m.replies, m.likes, m.views
        FROM posts p
        JOIN post_likes l ON p.id = l.post_id
        LEFT JOIN post_public_metrics m ON p.id = m.post_id
        WHERE l.user_id = $1 AND p.deleted_at IS NULL
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
		return nil, nil, fmt.Errorf("failed to list user likes: %w", err)
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
