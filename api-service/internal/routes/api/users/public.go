package users

import (
	"yabro.io/social-api/internal/db/userdb"
	"yabro.io/social-api/internal/dto"
	"yabro.io/social-api/internal/service"
	"yabro.io/social-api/internal/util"
)

func ToPublicUser(userData service.UserData) dto.User {
	user := userData.User
	return dto.User{
		ID:   util.Int64ToString(user.ID),
		Name: user.Name,

		Username:  user.Username,
		PfpURL:    user.PfpURL,
		Protected: user.Protected,
		CreatedAt: user.CreatedAt,

		Profile: toPublicProfile(userData.Profile),
	}
}

func toPublicProfile(profile *userdb.UserProfile) *dto.UserProfile {
	if profile == nil {
		return nil
	}

	return &dto.UserProfile{
		BannerUrl: profile.BannerUrl,
		Bio:       profile.Bio,

		Website:        profile.Website,
		Location:       profile.Location,
		Birthday:       profile.Birthday,
		PinnedPostID:   profile.PinnedPostID,
		PostCount:      profile.Posts,
		FollowerCount:  profile.Followers,
		FollowingCount: profile.Following,
	}
}
