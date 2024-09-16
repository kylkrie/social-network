package service

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"yabro.io/social-api/internal/db"
	"yabro.io/social-api/internal/db/userdb"
)

type UserService struct {
	userDB        *userdb.UserDB
	snowflakeNode *snowflake.Node
	minioStorage  *db.MinioStorage
}

func NewUserService(userDB *userdb.UserDB, snowflakeNode *snowflake.Node, minioStorage *db.MinioStorage) (*UserService, error) {
	return &UserService{
		userDB:        userDB,
		snowflakeNode: snowflakeNode,
		minioStorage:  minioStorage,
	}, nil
}

func (s *UserService) GetUserID(authUUID uuid.UUID) (int64, error) {
	return s.userDB.GetUserID(authUUID)
}

type UserData struct {
	User    userdb.User
	Profile *userdb.UserProfile
}

func (s *UserService) GetUserByID(id int64, includeProfile bool) (*UserData, error) {
	user, profile, err := s.userDB.GetUser(
		userdb.UserLookup{ID: &id},
		includeProfile,
	)
	if err != nil {
		return nil, err
	}

	return &UserData{
		User:    *user,
		Profile: profile,
	}, nil
}

func (s *UserService) GetUserByUsername(username string, includeProfile bool) (*UserData, error) {
	user, profile, err := s.userDB.GetUser(
		userdb.UserLookup{Username: &username},
		includeProfile,
	)
	if err != nil {
		return nil, err
	}

	return &UserData{
		User:    *user,
		Profile: profile,
	}, nil
}

func (s *UserService) CreateUser(authUUID uuid.UUID, name string, username string) (*UserData, error) {
	// Generate a new snowflake ID
	id := s.snowflakeNode.Generate().Int64()

	createParams := userdb.CreateUserParams{
		ID:       id,
		AuthUUID: authUUID,
		Name:     name,
		Username: username,
	}

	user, err := s.userDB.CreateUser(createParams)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &UserData{User: *user}, nil
}
