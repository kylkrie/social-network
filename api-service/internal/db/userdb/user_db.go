package userdb

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserDB struct {
	db *sqlx.DB
}

func NewUserDB(db *sqlx.DB) *UserDB {
	return &UserDB{db: db}
}

type User struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Username  string    `db:"username"`
	PfpURL    *string   `db:"pfp_url"`
	Protected bool      `db:"protected"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UserAuth struct {
	AuthUUID  uuid.UUID `db:"auth_uuid"`
	UserID    int64     `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}

type UserProfile struct {
	UserID       int64      `db:"user_id"`
	BannerUrl    *string    `db:"banner_url"`
	Bio          *string    `db:"bio"`
	Website      *string    `db:"website"`
	Location     *string    `db:"location"`
	Birthday     *time.Time `db:"birthday"`
	PinnedPostID *int64     `db:"pinned_post_id"`
	Posts        int        `db:"posts"`
	Followers    int        `db:"followers"`
	Following    int        `db:"following"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at"`
}
