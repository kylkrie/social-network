package app

import (
	"github.com/bwmarrin/snowflake"
	"github.com/jmoiron/sqlx"
	"yabro.io/social-api/db/userdb"
	"yabro.io/social-api/service"
)

type AppServices struct {
	UserService *service.UserService
}

func NewAppServices(db *sqlx.DB, snowflakeNode *snowflake.Node) (*AppServices, error) {
	userDb := userdb.NewUserDB(db)
	userService, err := service.NewUserService(userDb, snowflakeNode)
	if err != nil {
		return nil, err
	}
	services := &AppServices{
		UserService: userService,
	}

	return services, nil
}
