package postdb

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"yabro.io/social-api/internal/logger"
)

type CreatePostParams struct {
	ID             int64  `db:"id"`
	Content        string `db:"content"`
	AuthorID       int64  `db:"author_id"`
	ConversationID *int64 `db:"conversation_id"`
}

type Post struct {
	CreatePostParams
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

func (p Post) GetID() int64 { return p.ID }

func (pdb *PostDB) GetExecer(tx *sqlx.Tx) sqlx.ExtContext {
	var exec sqlx.ExtContext = pdb.db
	if tx != nil {
		exec = tx
	}

	return exec
}

func (pdb *PostDB) CreatePost(ctx context.Context, params CreatePostParams, tx *sqlx.Tx) error {
	exec := pdb.GetExecer(tx)
	query := `
		INSERT INTO posts (id, content, author_id, conversation_id)
		VALUES (:id, :content, :author_id, :conversation_id)
	`
	_, err := sqlx.NamedExecContext(ctx, exec, query, params)
	if err != nil {
		logger.Warn(ctx).Str("table", "posts").Err(err).Msg("CreatePost query failed")
		return err
	}
	return nil
}

func (pdb *PostDB) GetPost(ctx context.Context, postID int64) (*Post, error) {
	var post Post
	query := `
		SELECT id, content, author_id, conversation_id, created_at, updated_at, deleted_at
		FROM posts
		WHERE id = $1 AND deleted_at IS NULL
	`
	err := pdb.db.GetContext(ctx, &post, query, postID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Warn(ctx).Str("table", "posts").Err(err).Msg("GetPost query failed")
		return nil, err
	}
	return &post, nil
}

func (pdb *PostDB) GetPosts(ctx context.Context, postIDs []int64) ([]Post, error) {
	query := `
		SELECT id, content, author_id, conversation_id, created_at, updated_at, deleted_at
		FROM posts
		WHERE id = ANY($1) AND deleted_at IS NULL
	`
	var posts []Post
	err := pdb.db.SelectContext(ctx, &posts, query, postIDs)
	if err != nil {
		logger.Warn(ctx).Str("table", "posts").Err(err).Msg("GetPosts query failed")
		return nil, err
	}
	return posts, nil
}

type ListPostParams struct {
	UserID  *int64
	Limit   int
	Cursor  *int64
	IsReply bool
}

func (pdb *PostDB) ListPosts(ctx context.Context, p ListPostParams) ([]Post, *int64, error) {
	query := strings.Builder{}
	query.WriteString("SELECT * FROM posts WHERE deleted_at IS NULL")
	args := []interface{}{}

	if p.UserID != nil {
		query.WriteString(fmt.Sprintf(" AND author_id = $%d", len(args)+1))
		args = append(args, *p.UserID)
	}

	if p.IsReply {
		query.WriteString(" AND conversation_id IS NOT NULL")
	} else {
		query.WriteString(" AND conversation_id IS NULL")
	}

	if p.Cursor != nil {
		query.WriteString(fmt.Sprintf(" AND id < $%d", len(args)+1))
		args = append(args, *p.Cursor)
	}

	query.WriteString(" ORDER BY id DESC")
	query.WriteString(fmt.Sprintf(" LIMIT $%d", len(args)+1))
	args = append(args, p.Limit+1)

	rows, err := pdb.db.QueryxContext(ctx, query.String(), args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list posts: %w", err)
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.StructScan(&post); err != nil {
			return nil, nil, fmt.Errorf("failed to scan post: %w", err)
		}
		posts = append(posts, post)
	}

	var nextCursor *int64
	if len(posts) > p.Limit {
		nextCursor = &posts[p.Limit-1].ID
		posts = posts[:p.Limit]
	}

	return posts, nextCursor, nil
}

func (pdb *PostDB) DeletePost(ctx context.Context, id int64, authorID int64, tx *sqlx.Tx) error {
	exec := pdb.GetExecer(tx)
	query := `
		UPDATE posts
		SET deleted_at = CURRENT_TIMESTAMP
		WHERE id = $1 AND author_id = $2 AND deleted_at IS NULL
	`

	_, err := exec.ExecContext(ctx, query, id, authorID)
	if err != nil {
		logger.Warn(ctx).Str("table", "posts").Err(err).Msg("DeletePost query failed")
		return err
	}

	return nil
}
