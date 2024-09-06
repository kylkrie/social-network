package userdb

import (
	"fmt"
	"time"
)

type UpdateUserParams struct {
	Name      *string
	Username  *string
	PfpURL    *string
	Protected *bool
}

type UpdateUserProfileParams struct {
	Bio          *string
	Website      *string
	Location     *string
	Birthday     *time.Time
	PinnedPostID *int64
}

func (udb *UserDB) UpdateUser(userID int64, user *UpdateUserParams, profile *UpdateUserProfileParams) error {
	if user == nil && profile == nil {
		return nil
	}

	tx, err := udb.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	if user != nil {
		query := `
			UPDATE users
			SET
				name = COALESCE($2, name),
				username = COALESCE($3, username),
				pfp_url = COALESCE($4, pfp_url),
				protected = COALESCE($5, protected)
			WHERE id = $1
		`
		_, err = tx.Exec(query, userID, user.Name, user.Username, user.PfpURL, user.Protected)
		if err != nil {
			return fmt.Errorf("failed to update user: %w", err)
		}
	}

	if profile != nil {
		query := `
			UPDATE user_profiles
			SET
				bio = COALESCE($2, bio),
				website = COALESCE($3, website),
				location = COALESCE($4, location),
				birthday = COALESCE($5, birthday),
				pinned_post_id = COALESCE($6, pinned_post_id)
			WHERE user_id = $1
		`
		_, err = tx.Exec(query, userID, profile.Bio, profile.Website, profile.Location, profile.Birthday, profile.PinnedPostID)
		if err != nil {
			return fmt.Errorf("failed to update user profile: %w", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
