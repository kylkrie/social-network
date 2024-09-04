package profilestore

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ProfileStore struct {
	db *sqlx.DB
}

func NewProfileStore(db *sqlx.DB) *ProfileStore {
	return &ProfileStore{db: db}
}

func (s *ProfileStore) Create(params CreateProfileParams) (*Profile, error) {
	profile := &Profile{
		ID:        uuid.New(),
		UserID:    params.UserID,
		Username:  params.Username,
		Email:     params.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	query := `
		INSERT INTO profiles (id, user_id, username, email, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := s.db.Exec(query, profile.ID, profile.UserID, profile.Username, profile.Email, profile.CreatedAt, profile.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (s *ProfileStore) Get(id uuid.UUID) (*Profile, error) {
	var profile Profile
	err := s.db.Get(&profile, "SELECT * FROM profiles WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (s *ProfileStore) Update(id uuid.UUID, params UpdateProfileParams) (*Profile, error) {
	profile, err := s.Get(id)
	if err != nil {
		return nil, err
	}

	if params.Username != "" {
		profile.Username = params.Username
	}
	if params.Email != "" {
		profile.Email = params.Email
	}
	profile.UpdatedAt = time.Now()

	query := `
		UPDATE profiles
		SET username = $2, email = $3, updated_at = $4
		WHERE id = $1
	`
	_, err = s.db.Exec(query, profile.ID, profile.Username, profile.Email, profile.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (s *ProfileStore) Delete(id uuid.UUID) error {
	_, err := s.db.Exec("DELETE FROM profiles WHERE id = $1", id)
	return err
}

func (s *ProfileStore) List(cursor string, limit int) ([]Profile, string, error) {
	query := `
		SELECT * FROM profiles
		WHERE ($1 = '' OR created_at < (SELECT created_at FROM profiles WHERE id = $1::uuid))
		ORDER BY created_at DESC
		LIMIT $2
	`
	var profiles []Profile
	err := s.db.Select(&profiles, query, cursor, limit+1)
	if err != nil {
		return nil, "", err
	}

	var nextCursor string
	if len(profiles) > limit {
		nextCursor = profiles[len(profiles)-1].ID.String()
		profiles = profiles[:limit]
	}

	return profiles, nextCursor, nil
}
