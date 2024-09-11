package app

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bwmarrin/snowflake"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"yabro.io/social-api/internal/apperror"
	"yabro.io/social-api/internal/auth"
)

type AppState struct {
	AuthConfig *AuthConfig
	DB         *sqlx.DB
	JWKS       *auth.JWKS
	Services   *AppServices
	Validator  *validator.Validate
}

func CreateAppState() (*AppState, error) {
	// config
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// postgres
	dbpool, err := CreatePool()
	if err != nil {
		return nil, err
	}

	// jwks
	jwks_url := os.Getenv("JWKS_URL")
	if jwks_url == "" {
		return nil, fmt.Errorf("JWKS_URL environment variable not set")
	}

	jwks, err := auth.NewJWKS(jwks_url)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize JWKS: %v", err)
	}

	nodeIDStr := os.Getenv("NODE_ID")
	nodeID, err := strconv.ParseInt(nodeIDStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("NODE_ID environment variable not set (int64)")
	}

	snowflakeNode, err := snowflake.NewNode(nodeID)
	if err != nil {
		return nil, fmt.Errorf("failed to create snowflake node: %v", err)
	}

	minioClient, err := SetupMinioClient()
	services, err := NewAppServices(dbpool, snowflakeNode, minioClient)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize AppServices: %v", err)
	}

	validator := validator.New()
	apperror.SetupValidator(validator)

	if err != nil {
		return nil, fmt.Errorf("failed to setup Minio client: %v", err)
	}

	appState := &AppState{
		AuthConfig: cfg,
		DB:         dbpool,
		JWKS:       jwks,
		Services:   services,
		Validator:  validator,
	}

	return appState, nil
}

func (a *AppState) Close() {
	if a.DB != nil {
		a.DB.Close()
	}
}
