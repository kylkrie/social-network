package userdb

import (
	"context"
	"fmt"
	"strings"
)

func (udb *UserDB) GetMany(ctx context.Context, ids []int64) ([]User, error) {
	if len(ids) == 0 {
		return []User{}, nil
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
		FROM users
		WHERE id IN (%s)
	`, strings.Join(placeholders, ","))

	var users []User
	err := udb.db.SelectContext(ctx, &users, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	return users, nil
}
