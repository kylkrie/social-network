package dto

import (
	"yabro.io/social-api/internal/db/postdb"
	"yabro.io/social-api/internal/util"
)

type IncludeData struct {
	Posts            []Post                `json:"posts,omitempty"`
	Users            []User                `json:"users,omitempty"`
	UserInteractions []UserPostInteraction `json:"user_interactions,omitempty"`
	Media            []Media               `json:"media,omitempty"`
}

type UserPostInteraction struct {
	PostID       string `json:"post_id"`
	IsLiked      bool   `json:"is_liked"`
	IsBookmarked bool   `json:"is_bookmarked"`
}

func ToPublicUserPostInteractions(interaction postdb.UserPostInteraction) UserPostInteraction {
	return UserPostInteraction{
		PostID: util.Int64ToString(interaction.PostID),

		IsLiked:      interaction.IsLiked,
		IsBookmarked: interaction.IsBookmarked,
	}
}
