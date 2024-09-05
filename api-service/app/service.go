package app

import (
	"github.com/jmoiron/sqlx"
	userdb "yabro.io/social-api/db/user"
	"yabro.io/social-api/service"
)

type AppServices struct {
	UserService *service.UserService
}

func NewAppServices(db *sqlx.DB, nodeID int64) (*AppServices, error) {
	userDb := userdb.NewUserDB(db)
	userService, err := service.NewUserService(userDb, nodeID)
	if err != nil {
		return nil, err
	}
	services := &AppServices{
		UserService: userService,
	}

	return services, nil
}
