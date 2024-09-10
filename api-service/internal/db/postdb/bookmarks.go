package postdb

import (
	"fmt"
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
