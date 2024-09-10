package dto

type IncludeData struct {
	Posts            *[]Post                `json:"posts,omitempty"`
	Users            *[]User                `json:"users,omitempty"`
	UserInteractions *[]UserPostInteraction `json:"user_interactions,omitempty"`
}

type UserPostInteraction struct {
	PostID       string `json:"post_id"`
	IsLiked      bool   `json:"is_liked"`
	IsBookmarked bool   `json:"is_bookmarked"`
}
