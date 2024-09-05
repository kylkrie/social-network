package userdb

import (
	"github.com/google/uuid"
)

func (s *UserDB) GetUserID(authUUID uuid.UUID) (int64, error) {
	var userAuth UserAuth
	err := s.db.Get(&userAuth, "SELECT * FROM user_auth WHERE auth_uuid = $1", authUUID)
	if err != nil {
		return -1, err
	}
	return userAuth.UserID, nil
}
