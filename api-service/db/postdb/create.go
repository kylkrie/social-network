package postdb

import (
	"database/sql"
	"fmt"
)

type CreatePostParams struct {
	ID             int64
	Content        string
	AuthorID       int64
	ConversationID *int64
}

func (pdb *PostDB) CreatePost(p CreatePostParams) (*Post, error) {
	query := `
		INSERT INTO posts (id, content, author_id, conversation_id)
		VALUES ($1, $2, $3, $4)
		RETURNING *
	`

	var post Post
	err := pdb.db.QueryRowx(query, p.ID, p.Content, p.AuthorID, p.ConversationID).StructScan(&post)
	if err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	return &post, nil
}

func (pdb *PostDB) GetPost(id int64, includeMetrics bool) (*Post, *PostPublicMetrics, error) {
	query := `SELECT * FROM posts WHERE id = $1 AND deleted_at IS NULL`

	var post Post
	err := pdb.db.Get(&post, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, fmt.Errorf("post not found")
		}
		return nil, nil, fmt.Errorf("failed to get post: %w", err)
	}

	var metrics *PostPublicMetrics
	if includeMetrics {
		metricsQuery := `SELECT * FROM post_public_metrics WHERE post_id = $1`
		metrics = &PostPublicMetrics{}
		err = pdb.db.Get(metrics, metricsQuery, id)
		if err != nil && err != sql.ErrNoRows {
			return nil, nil, fmt.Errorf("failed to get post metrics: %w", err)
		}
	}

	return &post, metrics, nil
}

type UpdatePostParams struct {
	ID      int64
	Content string
}

func (pdb *PostDB) UpdatePost(p UpdatePostParams) (*Post, error) {
	query := `
		UPDATE posts
		SET content = $2, updated_at = CURRENT_TIMESTAMP
		WHERE id = $1 AND deleted_at IS NULL
		RETURNING *
	`

	var post Post
	err := pdb.db.QueryRowx(query, p.ID, p.Content).StructScan(&post)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("post not found")
		}
		return nil, fmt.Errorf("failed to update post: %w", err)
	}

	return &post, nil
}

func (pdb *PostDB) DeletePost(id int64, authorID int64) error {
	query := `
		UPDATE posts
		SET deleted_at = CURRENT_TIMESTAMP
		WHERE id = $1 AND author_id = $2 AND deleted_at IS NULL
	`

	_, err := pdb.db.Exec(query, id, authorID)
	if err != nil {
		return fmt.Errorf("failed to soft delete post: %w", err)
	}

	return nil
}

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
