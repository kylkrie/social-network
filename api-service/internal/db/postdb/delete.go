package postdb

import "fmt"

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
