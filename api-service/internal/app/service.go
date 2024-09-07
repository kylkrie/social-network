package app

import (
	"github.com/bwmarrin/snowflake"
	"github.com/jmoiron/sqlx"
	"yabro.io/social-api/internal/db/postdb"
	"yabro.io/social-api/internal/db/userdb"
	"yabro.io/social-api/internal/service"
)

type AppServices struct {
	UserService    *service.UserService
	PostService    *service.PostService
	IncludeService *service.IncludeService
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

	includeService := service.NewIncludeService(userDb, postDb)

	services := &AppServices{
		UserService:    userService,
		PostService:    postService,
		IncludeService: includeService,
	}

	return services, nil
}
