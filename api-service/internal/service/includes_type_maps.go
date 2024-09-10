package service

import (
	"yabro.io/social-api/internal/db/postdb"
	"yabro.io/social-api/internal/dto"
	"yabro.io/social-api/internal/util"
)

func toPublicUserPostInteractions(interaction *postdb.UserPostInteraction) *dto.UserPostInteraction {
	if interaction == nil {
		return nil
	}
	return &dto.UserPostInteraction{
		PostID:       util.Int64ToString(interaction.PostID),
		IsLiked:      interaction.IsLiked,
		IsBookmarked: interaction.IsBookmarked,
	}
}
