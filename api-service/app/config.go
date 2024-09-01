package app

import (
	"os"

	"golang.org/x/oauth2"
)

type Config struct {
	OAuthConfig   *oauth2.Config
	AllowedOrigin string
	JWKS_URL      string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		OAuthConfig: &oauth2.Config{
			ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
			ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
			RedirectURL:  os.Getenv("OAUTH_REDIRECT_URI"),
			Endpoint: oauth2.Endpoint{
				AuthURL:  os.Getenv("OAUTH_AUTH_URL"),
				TokenURL: os.Getenv("OAUTH_TOKEN_URL"),
			},
			Scopes: []string{"email", "profile"},
		},
		AllowedOrigin: os.Getenv("ALLOWED_ORIGIN"),
	}

	// Add any necessary validation here

	return cfg, nil
}
