package service

import (
	"fmt"
	"time"

	"yabro.io/social-api/internal/db/userdb"
)

type UpdateUserParams struct {
	Name         *string
	Protected    *bool
	Bio          *string
	Website      *string
	Location     *string
	Birthday     *time.Time
	PinnedPostID *int64
}

func (s *UserService) UpdateUser(userID int64, u *UpdateUserParams) error {
	var updateUser *userdb.UpdateUserParams
	var updateProfile *userdb.UpdateUserProfileParams

	if u.Name != nil || u.Protected != nil {
		updateUser = &userdb.UpdateUserParams{
			Name:      u.Name,
			Protected: u.Protected,
		}
	}
	if u.Bio != nil || u.Website != nil || u.Location != nil || u.Birthday != nil || u.PinnedPostID != nil {
		updateProfile = &userdb.UpdateUserProfileParams{
			Bio:          u.Bio,
			Website:      u.Website,
			Location:     u.Location,
			Birthday:     u.Birthday,
			PinnedPostID: u.PinnedPostID,
		}
	}

	// Perform the update
	err := s.userDB.UpdateUser(userID, updateUser, updateProfile)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
