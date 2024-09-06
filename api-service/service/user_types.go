package service

import (
	"time"

	"yabro.io/social-api/db/userdb"
)

type PublicUser struct {
	ID        int64              `json:"id"`
	Name      string             `json:"name"`
	Username  string             `json:"username"`
	PfpURL    *string            `json:"pfp_url"`
	Protected bool               `json:"protected"`
	Profile   *PublicUserProfile `json:"profile,omitempty"`
}

type PublicUserProfile struct {
	BannerUrl      *string    `json:"banner_url"`
	Bio            *string    `json:"bio"`
	Website        *string    `json:"website"`
	Location       *string    `json:"location"`
	Birthday       *time.Time `json:"birthday"`
	PinnedPostID   *int64     `json:"pinned_post_id"`
	FollowerCount  int        `json:"follower_count"`
	FollowingCount int        `json:"following_count"`
}

func toPublicUser(user *userdb.User, profile *userdb.UserProfile) PublicUser {
	return PublicUser{
		ID:   user.ID,
		Name: user.Name,

		Username:  user.Username,
		PfpURL:    user.PfpURL,
		Protected: user.Protected,

		Profile: toPublicProfile(profile),
	}
}

func toPublicProfile(profile *userdb.UserProfile) *PublicUserProfile {
	if profile == nil {
		return nil
	}

	return &PublicUserProfile{
		BannerUrl: profile.BannerUrl,
		Bio:       profile.Bio,

		Website:      profile.Website,
		Location:     profile.Location,
		Birthday:     profile.Birthday,
		PinnedPostID: profile.PinnedPostID,
	}
}
