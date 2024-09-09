package dto

import (
	"time"
)

type User struct {
	ID        string       `json:"id"`
	Name      string       `json:"name"`
	Username  string       `json:"username"`
	PfpURL    *string      `json:"pfp_url"`
	Protected bool         `json:"protected"`
	CreatedAt time.Time    `json:"created_at"`
	Profile   *UserProfile `json:"profile,omitempty"`
}

type UserProfile struct {
	BannerUrl      *string    `json:"banner_url"`
	Bio            *string    `json:"bio"`
	Website        *string    `json:"website"`
	Location       *string    `json:"location"`
	Birthday       *time.Time `json:"birthday"`
	PinnedPostID   *int64     `json:"pinned_post_id"`
	PostCount      int        `json:"post_count"`
	FollowerCount  int        `json:"follower_count"`
	FollowingCount int        `json:"following_count"`
}
