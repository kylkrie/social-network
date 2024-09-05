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
	Bio          *string    `db:"bio"`
	Website      *string    `db:"website"`
	Location     *string    `db:"location"`
	Birthday     *time.Time `db:"birthday"`
	PinnedPostID *int64     `db:"pinned_post_id"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at"`
}
type UserPublicMetrics struct {
	UserID         int64     `db:"user_id"`
	FollowersCount int       `db:"followers_count"`
	FollowingCount int       `db:"following_count"`
	PostCount      int       `db:"post_count"`
	ListedCount    int       `db:"listed_count"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
