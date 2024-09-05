package userdb

type UpdateUserPublicMetricsParams struct {
	FollowersCount *int
	FollowingCount *int
	PostCount      *int
	ListedCount    *int
}

func (udb *UserDB) UpdateUserPublicMetrics(userID int64, params UpdateUserPublicMetricsParams) (*UserPublicMetrics, error) {
	query := `
		UPDATE user_public_metrics
		SET
			followers_count = COALESCE($2, followers_count),
			following_count = COALESCE($3, following_count),
			post_count = COALESCE($4, post_count),

			listed_count = COALESCE($5, listed_count)
		WHERE user_id = $1
		RETURNING *
	`

	var metrics UserPublicMetrics
	err := udb.db.QueryRowx(query,
		userID,
		params.FollowersCount,
		params.FollowingCount,
		params.PostCount,
		params.ListedCount,
	).StructScan(&metrics)
	if err != nil {
		return nil, err
	}

	return &metrics, nil
}

func (udb *UserDB) GetUserPublicMetrics(userID int64) (*UserPublicMetrics, error) {
	query := `
		SELECT *
		FROM user_public_metrics
		WHERE user_id = $1

	`

	var metrics UserPublicMetrics
	err := udb.db.QueryRowx(query, userID).StructScan(&metrics)
	if err != nil {
		return nil, err
	}

	return &metrics, nil
}
