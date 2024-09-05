package service

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"yabro.io/social-api/db/userdb"
)

type UserService struct {
	userDB        *userdb.UserDB
	snowflakeNode *snowflake.Node
}

func NewUserService(userDB *userdb.UserDB, snowflakeNode *snowflake.Node) (*UserService, error) {
	return &UserService{
		userDB:        userDB,
		snowflakeNode: snowflakeNode,
	}, nil
}

func (s *UserService) GetUserID(authUUID uuid.UUID) (int64, error) {
	return s.userDB.GetUserID(authUUID)
}

func (s *UserService) GetUserByID(id int64, includeProfile bool, includeMetrics bool) (*PublicUser, error) {
	user, profile, metrics, err := s.userDB.GetUser(
		userdb.UserLookup{ID: &id},
		includeProfile,
		includeMetrics,
	)
	if err != nil {
		return nil, err
	}

	publicUser := toPublicUser(user, profile, metrics)
	return &publicUser, nil
}

func (s *UserService) GetUserByUsername(username string, includeProfile bool, includeMetrics bool) (*PublicUser, error) {
	user, profile, metrics, err := s.userDB.GetUser(
		userdb.UserLookup{Username: &username},
		includeProfile,
		includeMetrics,
	)
	if err != nil {
		return nil, err
	}

	publicUser := toPublicUser(user, profile, metrics)
	return &publicUser, nil
}

func (s *UserService) CreateUser(authUUID uuid.UUID, name string, username string) (*userdb.User, error) {
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

	return user, nil
}
