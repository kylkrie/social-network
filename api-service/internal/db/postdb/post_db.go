package postdb

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type PostDB struct {
	db *sqlx.DB
}

func NewPostDB(db *sqlx.DB) *PostDB {
	return &PostDB{db: db}
}

type PostWithMetrics struct {
	Post    Post
	Metrics PostPublicMetrics
}

func (pdb *PostDB) BeginTx(ctx context.Context) (*sqlx.Tx, error) {
	return pdb.db.BeginTxx(ctx, nil)
}
