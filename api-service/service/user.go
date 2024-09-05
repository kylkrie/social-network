package service

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	userdb "yabro.io/social-api/db/user"
)

type UserService struct {
	userDB      *userdb.UserDB
	idGenerator *snowflake.Node
}

func NewUserService(userDB *userdb.UserDB, nodeID int64) (*UserService, error) {
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		return nil, fmt.Errorf("failed to create snowflake node: %w", err)
	}

	return &UserService{
		userDB:      userDB,
		idGenerator: node,
	}, nil
}

type CreateUserInput struct {
	AuthUUID uuid.UUID
	Name     string
	Username string
}

func (s *UserService) GetUserID(authUUID uuid.UUID) (int64, error) {
	return s.userDB.GetUserID(authUUID)
}

func (s *UserService) GetUser(id int64) (*userdb.User, error) {
	return s.userDB.GetUser(id)
}

func (s *UserService) CreateUser(input CreateUserInput) (*userdb.User, error) {
	// Generate a new snowflake ID
	id := s.idGenerator.Generate().Int64()

	// Prepare the parameters for UserDB.CreateUser
	createParams := userdb.CreateUserParams{
		ID:       id,
		AuthUUID: input.AuthUUID,
		Name:     input.Name,
		Username: input.Username,
	}

	// Call UserDB.CreateUser
	user, err := s.userDB.CreateUser(createParams)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}
