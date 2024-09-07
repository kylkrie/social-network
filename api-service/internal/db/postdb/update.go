package postdb

import (
	"database/sql"
	"fmt"
)

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
