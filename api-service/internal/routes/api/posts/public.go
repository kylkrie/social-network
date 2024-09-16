package posts

import (
	"yabro.io/social-api/internal/dto"
	"yabro.io/social-api/internal/service"
	"yabro.io/social-api/internal/util"
)

type PostListResponse struct {
	Data       []dto.Post      `json:"data"`
	Includes   dto.IncludeData `json:"includes"`
	NextCursor *string         `json:"next_cursor,omitempty"`
}

type PostResponse struct {
	Data     dto.Post        `json:"data"`
	Includes dto.IncludeData `json:"includes"`
}

func ToPostResponse(postData service.PostData, includeData service.IncludeData) PostResponse {
	data := dto.ToPublicPost(postData, includeData)

	includes := ToPublicIncludes(includeData)

	return PostResponse{
		Data:     data,
		Includes: includes,
	}
}

func ToPostListResponse(postData []service.PostData, includeData service.IncludeData, nextCursor *int64) PostListResponse {
	data := make([]dto.Post, len(postData))
	for i, pd := range postData {
		data[i] = dto.ToPublicPost(pd, includeData)
	}

	includes := ToPublicIncludes(includeData)

	return PostListResponse{
		Data:       data,
		Includes:   includes,
		NextCursor: util.NullableInt64ToString(nextCursor),
	}
}

func ToPublicIncludes(includeData service.IncludeData) dto.IncludeData {
	dtoPosts := make([]dto.Post, len(includeData.Posts))
	for i, p := range includeData.Posts {
		dtoPosts[i] = dto.ToPublicPost(service.PostData{Post: p}, includeData)
	}

	dtoUsers := make([]dto.User, len(includeData.Users))
	for i, u := range includeData.Users {
		dtoUsers[i] = dto.ToPublicUser(&u, nil)
	}

	var dtoInteractions []dto.UserPostInteraction
	for _, upi := range includeData.Interactions {
		dtoInteractions = append(dtoInteractions, dto.ToPublicUserPostInteractions(upi))
	}

	var dtoMedia []dto.Media
	for _, mediaList := range includeData.Media {
		for _, m := range mediaList {
			dtoMedia = append(dtoMedia, dto.ToPublicMedia(m))
		}
	}

	return dto.IncludeData{
		Posts:            dtoPosts,
		Users:            dtoUsers,
		UserInteractions: dtoInteractions,
		Media:            dtoMedia,
	}
}
