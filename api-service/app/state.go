package app

import (
	"fmt"
	"log"
	"os"

	"yabro.io/social-api/auth"
	"yabro.io/social-api/postgres"

	"github.com/jmoiron/sqlx"
)

type AppState struct {
	Config *Config
	DB     *sqlx.DB
	JWKS   *auth.JWKS
	Stores *AppStores
}

func CreateAppState() (*AppState, error) {
	// config
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// postgres
	dbpool, err := postgres.CreatePool()
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

	appState := &AppState{
		Config: cfg,
		DB:     dbpool,
		JWKS:   jwks,
		Stores: CreateStores(dbpool),
	}

	return appState, nil
}

func (a *AppState) Close() {
	if a.DB != nil {
		a.DB.Close()
	}
}
