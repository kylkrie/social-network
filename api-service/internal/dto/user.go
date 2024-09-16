package dto

import (
	"time"

	"yabro.io/social-api/internal/db/userdb"
	"yabro.io/social-api/internal/util"
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

func ToPublicUser(user *userdb.User, profile *userdb.UserProfile) User {
	return User{
		ID:   util.Int64ToString(user.ID),
		Name: user.Name,

		Username:  user.Username,
		PfpURL:    user.PfpURL,
		Protected: user.Protected,
		CreatedAt: user.CreatedAt,

		Profile: toPublicProfile(profile),
	}
}

func toPublicProfile(profile *userdb.UserProfile) *UserProfile {
	if profile == nil {
		return nil
	}

	return &UserProfile{
		BannerUrl: profile.BannerUrl,

		Bio: profile.Bio,

		Website:        profile.Website,
		Location:       profile.Location,
		Birthday:       profile.Birthday,
		PinnedPostID:   profile.PinnedPostID,
		PostCount:      profile.Posts,
		FollowerCount:  profile.Followers,
		FollowingCount: profile.Following,
	}
}
