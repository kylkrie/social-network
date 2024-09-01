package services

import (
	"context"

	"golang.org/x/oauth2"
)

func ExchangeCodeForToken(ctx context.Context, config *oauth2.Config, code string) (*oauth2.Token, error) {
	return config.Exchange(ctx, code)
}
