package auth

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/lestrrat-go/jwx/jwk"
)

const (
	maxRetries = 5
	retryDelay = 10 * time.Second
)

type JWKS struct {
	keyset    jwk.Set
	mu        sync.RWMutex
	lastFetch time.Time
}

func NewJWKS(jwksURL string) (*JWKS, error) {
	k := &JWKS{}
	err := k.updateKeysWithRetry(jwksURL)
	if err != nil {
		return nil, err
	}
	return k, nil
}

func (k *JWKS) updateKeysWithRetry(jwksURL string) error {
	var err error
	for attempt := 1; attempt <= maxRetries; attempt++ {
		err = k.updateKeys(jwksURL)
		if err == nil {
			return nil
		}

		fmt.Printf("Attempt %d failed to fetch JWKS: %v. Retrying in %v...\n", attempt, err, retryDelay)
		time.Sleep(retryDelay)
	}
	return fmt.Errorf("failed to fetch JWKS after %d attempts: %v", maxRetries, err)
}

func (k *JWKS) updateKeys(jwksURL string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	keyset, err := jwk.Fetch(ctx, jwksURL)
	if err != nil {
		return fmt.Errorf("failed to fetch JWKS: %v", err)
	}

	k.mu.Lock()
	defer k.mu.Unlock()
	k.keyset = keyset
	k.lastFetch = time.Now()
	return nil
}

func (k *JWKS) getKey(token *jwt.Token) (interface{}, error) {
	k.mu.RLock()
	defer k.mu.RLock()

	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("expecting JWT header to have string kid")
	}

	key, found := k.keyset.LookupKeyID(keyID)
	if !found {
		return nil, fmt.Errorf("unable to find key %v", keyID)
	}

	var rawKey interface{}
	if err := key.Raw(&rawKey); err != nil {
		return nil, fmt.Errorf("unable to get raw key: %v", err)
	}

	return rawKey, nil
}
