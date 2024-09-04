package app

import (
	"github.com/jmoiron/sqlx"
	"yabro.io/social-api/stores/profilestore"
)

type AppStores struct {
	Profile *profilestore.ProfileStore
}

func CreateStores(db *sqlx.DB) *AppStores {
	appStores := &AppStores{
		Profile: profilestore.NewProfileStore(db),
	}

	return appStores
}
