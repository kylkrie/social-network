package postdb

import (
	"fmt"
	"strings"
)

func (pdb *PostDB) BookmarkPost(postID, userID int64) error {
	query := `

        INSERT INTO post_bookmarks (post_id, user_id)
        VALUES ($1, $2)
        ON CONFLICT (post_id, user_id) DO NOTHING
    `
	_, err := pdb.db.Exec(query, postID, userID)
	if err != nil {
		return fmt.Errorf("failed to bookmark post: %w", err)
	}
	return nil
}

func (pdb *PostDB) UnbookmarkPost(postID, userID int64) error {
	query := `
        DELETE FROM post_bookmarks
        WHERE post_id = $1 AND user_id = $2

    `

	_, err := pdb.db.Exec(query, postID, userID)
	if err != nil {
		return fmt.Errorf("failed to unbookmark post: %w", err)
	}
	return nil
}

func (pdb *PostDB) ListUserBookmarks(userID int64, limit int, cursor *int64) ([]PostData, *int64, error) {
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

	rows, err := pdb.db.Queryx(query.String(), args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list user bookmarks: %w", err)
	}
	defer rows.Close()

	var posts []PostData
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
		posts = append(posts, PostData{Post: post, Metrics: &metrics})
	}

	var nextCursor *int64
	if len(posts) > limit {

		nextCursor = &posts[limit-1].Post.ID
		posts = posts[:limit]
	}

	return posts, nextCursor, nil
}
