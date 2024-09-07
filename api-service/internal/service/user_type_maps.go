package service

import (
	"yabro.io/social-api/internal/db/userdb"
	"yabro.io/social-api/internal/dto"
)

func toPublicUser(user *userdb.User, profile *userdb.UserProfile) dto.User {
	return dto.User{
		ID:   user.ID,
		Name: user.Name,

		Username:  user.Username,
		PfpURL:    user.PfpURL,
		Protected: user.Protected,

		Profile: toPublicProfile(profile),
	}
}

func toPublicProfile(profile *userdb.UserProfile) *dto.UserProfile {
	if profile == nil {
		return nil
	}

	return &dto.UserProfile{
		BannerUrl: profile.BannerUrl,
		Bio:       profile.Bio,

		Website:      profile.Website,
		Location:     profile.Location,
		Birthday:     profile.Birthday,
		PinnedPostID: profile.PinnedPostID,
	}
}
