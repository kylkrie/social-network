package app

import (
	"github.com/bwmarrin/snowflake"
	"github.com/jmoiron/sqlx"
	"yabro.io/social-api/db/postdb"
	"yabro.io/social-api/db/userdb"
	"yabro.io/social-api/service"
)

type AppServices struct {
	UserService *service.UserService
	PostService *service.PostService
}

func NewAppServices(db *sqlx.DB, snowflakeNode *snowflake.Node) (*AppServices, error) {
	userDb := userdb.NewUserDB(db)
	userService, err := service.NewUserService(userDb, snowflakeNode)
	if err != nil {
		return nil, err
	}

	postDb := postdb.NewPostDB(db)
	postService, err := service.NewPostService(postDb, snowflakeNode)
	if err != nil {
		return nil, err
	}

	services := &AppServices{
		UserService: userService,
		PostService: postService,
	}

	return services, nil
}
