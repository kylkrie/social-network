package dto

type IncludeData struct {
	Posts *[]Post `json:"posts,omitempty"`
	Users *[]User `json:"users,omitempty"`
}
