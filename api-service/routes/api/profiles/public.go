package profiles

import "yabro.io/social-api/stores/profilestore"

type PublicProfile struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func FromEntity(p *profilestore.Profile) PublicProfile {
	return PublicProfile{
		UserID:   p.UserID,
		Username: p.Username,
		Email:    p.Email,
	}
}
