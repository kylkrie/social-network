package userdb

import (
	"time"
)

type UpdateUserProfileParams struct {
	Bio *string

	Website      *string
	Location     *string
	Birthday     *time.Time
	PinnedPostID *int64
}

func (udb *UserDB) UpdateUserProfile(userID int64, params UpdateUserProfileParams) (*UserProfile, error) {
	query := `
		UPDATE user_profiles
		SET
			bio = COALESCE($2, bio),
			website = COALESCE($3, website),
			location = COALESCE($4, location),
			birthday = COALESCE($5, birthday),
			pinned_post_id = COALESCE($6, pinned_post_id)
		WHERE user_id = $1
		RETURNING *
	`

	var profile UserProfile
	err := udb.db.QueryRowx(query,
		userID,
		params.Bio,
		params.Website,
		params.Location,
		params.Birthday,
		params.PinnedPostID,
	).StructScan(&profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (udb *UserDB) GetUserProfile(userID int64) (*UserProfile, error) {
	query := `
		SELECT *
		FROM user_profiles
		WHERE user_id = $1
	`

	var profile UserProfile
	err := udb.db.QueryRowx(query, userID).StructScan(&profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}
