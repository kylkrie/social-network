package app

import (
	"os"

	"github.com/bwmarrin/snowflake"
	"github.com/jmoiron/sqlx"
	"github.com/minio/minio-go/v7"
	"yabro.io/social-api/internal/db"
	"yabro.io/social-api/internal/db/postdb"
	"yabro.io/social-api/internal/db/userdb"
	"yabro.io/social-api/internal/service"
)

type AppServices struct {
	UserService    *service.UserService
	PostService    *service.PostService
	IncludeService *service.IncludeService
}

func NewAppServices(sqlxDB *sqlx.DB, snowflakeNode *snowflake.Node, minioClient *minio.Client) (*AppServices, error) {
	userDb := userdb.NewUserDB(sqlxDB)
	postDb := postdb.NewPostDB(sqlxDB)
	cdnBaseUrl := os.Getenv("CDN_BASE_URL")
	minioStorage := db.NewMinioStorage(minioClient, cdnBaseUrl)

	userService, err := service.NewUserService(userDb, snowflakeNode, minioStorage)
	if err != nil {
		return nil, err
	}

	postService, err := service.NewPostService(postDb, snowflakeNode, minioStorage)
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
