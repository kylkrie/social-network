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
	Metrics   *PublicUserMetrics `json:"metrics,omitempty"`
}

type PublicUserProfile struct {
	Bio          *string    `json:"bio"`
	Website      *string    `json:"website"`
	Location     *string    `json:"location"`
	Birthday     *time.Time `json:"birthday"`
	PinnedPostID *int64     `json:"pinned_post_id"`
}

type PublicUserMetrics struct {
	FollowersCount int `db:"followers_count"`
	FollowingCount int `db:"following_count"`
	PostCount      int `db:"post_count"`
	ListedCount    int `db:"listed_count"`
}

func toPublicUser(user *userdb.User, profile *userdb.UserProfile, metrics *userdb.UserPublicMetrics) PublicUser {
	return PublicUser{
		ID:   user.ID,
		Name: user.Name,

		Username:  user.Username,
		PfpURL:    user.PfpURL,
		Protected: user.Protected,

		Profile: toPublicProfile(profile),
		Metrics: toPublicMetrics(metrics),
	}
}

func toPublicProfile(profile *userdb.UserProfile) *PublicUserProfile {
	if profile == nil {
		return nil
	}

	return &PublicUserProfile{
		Bio: profile.Bio,

		Website:      profile.Website,
		Location:     profile.Location,
		Birthday:     profile.Birthday,
		PinnedPostID: profile.PinnedPostID,
	}
}

func toPublicMetrics(metrics *userdb.UserPublicMetrics) *PublicUserMetrics {
	if metrics == nil {
		return nil
	}

	return &PublicUserMetrics{
		FollowersCount: metrics.FollowersCount,
		FollowingCount: metrics.FollowingCount,

		PostCount:   metrics.PostCount,
		ListedCount: metrics.ListedCount,
	}
}
