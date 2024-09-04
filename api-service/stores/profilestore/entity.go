package profilestore

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ID        uuid.UUID `db:"id"`
	UserID    string    `db:"user_id"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type CreateProfileParams struct {
	UserID   string
	Username string
	Email    string
}

type UpdateProfileParams struct {
	Username string
	Email    string
}
