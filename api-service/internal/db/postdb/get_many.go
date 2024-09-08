package postdb

import (
	"fmt"
	"strings"
)

func (pdb *PostDB) GetMany(ids []int64) ([]*Post, error) {
	if len(ids) == 0 {
		return []*Post{}, nil
	}

	// Create a string of placeholders for the SQL query
	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = id
	}

	query := fmt.Sprintf(`
		SELECT *
		FROM posts
		WHERE id IN (%s) AND deleted_at IS NULL
	`, strings.Join(placeholders, ","))

	var posts []*Post
	err := pdb.db.Select(&posts, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get posts: %w", err)
	}

	return posts, nil
}
