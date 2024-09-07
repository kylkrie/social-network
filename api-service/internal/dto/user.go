package dto

import (
	"time"
)

type User struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	Username  string       `json:"username"`
	PfpURL    *string      `json:"pfp_url"`
	Protected bool         `json:"protected"`
	Profile   *UserProfile `json:"profile,omitempty"`
}

type UserProfile struct {
	BannerUrl      *string    `json:"banner_url"`
	Bio            *string    `json:"bio"`
	Website        *string    `json:"website"`
	Location       *string    `json:"location"`
	Birthday       *time.Time `json:"birthday"`
	PinnedPostID   *int64     `json:"pinned_post_id"`
	FollowerCount  int        `json:"follower_count"`
	FollowingCount int        `json:"following_count"`
}
