package profiles

import (
	"github.com/google/uuid"
	"yabro.io/social-api/stores/profilestore"
)

type PublicProfile struct {
	ID       uuid.UUID `json:"id"`
	UserID   string    `json:"user_id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}

func FromEntity(p *profilestore.Profile) PublicProfile {
	return PublicProfile{
		ID:       p.ID,
		UserID:   p.UserID,
		Username: p.Username,
		Email:    p.Email,
	}
}
