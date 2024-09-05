package app

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"yabro.io/social-api/auth"
)

type AppState struct {
	AuthConfig *AuthConfig
	DB         *sqlx.DB
	JWKS       *auth.JWKS
	Services   *AppServices
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

	services, err := NewAppServices(dbpool, 100)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize AppServices: %v", err)
	}

	appState := &AppState{
		AuthConfig: cfg,
		DB:         dbpool,
		JWKS:       jwks,
		Services:   services,
	}

	return appState, nil
}

func (a *AppState) Close() {
	if a.DB != nil {
		a.DB.Close()
	}
}
