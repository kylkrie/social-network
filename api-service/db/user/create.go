package userdb

import (
	"github.com/google/uuid"
)

type CreateUserParams struct {
	AuthUUID uuid.UUID
	ID       int64
	Name     string
	Username string
}

func (udb *UserDB) CreateUser(p CreateUserParams) (*User, error) {
	query := `
		WITH auth_insert AS (
			INSERT INTO user_auth (auth_uuid, user_id)
			VALUES ($1, $2)
		), profile_insert AS (
			INSERT INTO user_profiles (user_id)
			VALUES ($2)
		), metrics_insert AS (
			INSERT INTO user_public_metrics (user_id)
			VALUES ($2)
		)
		INSERT INTO users (id, name, username)
		VALUES ($2, $3, $4)
		RETURNING *
	`

	var user User
	err := udb.db.QueryRowx(query,
		p.AuthUUID, p.ID, p.Name, p.Username,
	).StructScan(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
