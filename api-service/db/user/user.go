package userdb

type UpdateUserParams struct {
	Name      *string
	Username  *string
	PfpURL    *string
	Protected *bool
}

func (udb *UserDB) UpdateUser(userID int64, params UpdateUserParams) (*User, error) {
	query := `
		UPDATE users
		SET
			name = COALESCE($2, name),
			username = COALESCE($3, username),
			pfp_url = COALESCE($4, pfp_url),
			protected = COALESCE($5, protected)
		WHERE id = $1
		RETURNING *
	`

	var user User
	err := udb.db.QueryRowx(query,

		userID,
		params.Name,

		params.Username,
		params.PfpURL,
		params.Protected,
	).StructScan(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (udb *UserDB) GetUser(userID int64) (*User, error) {
	query := `
		SELECT *
		FROM users
		WHERE id = $1
	`

	var user User
	err := udb.db.QueryRowx(query, userID).StructScan(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
