package postdb

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type ListPostParams struct {
	UserID         *int64
	Limit          int
	Cursor         *int64
	IsReply        bool
	ConversationID *int64
}

func (pdb *PostDB) ListPosts(p ListPostParams) ([]Post, *int64, error) {
	log.Info().Any("p", p).Msg("List")

	query := strings.Builder{}
	query.WriteString("SELECT * FROM posts WHERE deleted_at IS NULL")
	args := []interface{}{}

	if p.UserID != nil {
		query.WriteString(fmt.Sprintf(" AND author_id = $%d", len(args)+1))
		args = append(args, *p.UserID)
	}

	if p.ConversationID != nil {
		query.WriteString(fmt.Sprintf(" AND conversation_id = $%d", len(args)+1))
		args = append(args, *p.ConversationID)
	} else if p.IsReply {
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

	rows, err := pdb.db.Queryx(query.String(), args...)
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

func (pdb *PostDB) ListPostDatas(p ListPostParams) ([]PostData, *int64, error) {
	// First, fetch the posts
	posts, nextCursor, err := pdb.ListPosts(p)
	if err != nil {
		return nil, nil, err
	}

	if len(posts) == 0 {
		return []PostData{}, nextCursor, nil
	}

	// Collect post IDs
	postIDs := make([]int64, len(posts))
	for i, post := range posts {
		postIDs[i] = post.ID
	}

	// Fetch metrics
	metricsQuery := `SELECT * FROM post_public_metrics WHERE post_id IN (?)`
	query, args, err := sqlx.In(metricsQuery, postIDs)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create IN query for metrics: %w", err)
	}
	query = pdb.db.Rebind(query)

	var metrics []PostPublicMetrics
	err = pdb.db.Select(&metrics, query, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch metrics: %w", err)
	}

	// Fetch references
	referencesQuery := `SELECT * FROM post_references WHERE source_post_id IN (?)`
	query, args, err = sqlx.In(referencesQuery, postIDs)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create IN query for references: %w", err)
	}
	query = pdb.db.Rebind(query)

	var references []PostReference
	err = pdb.db.Select(&references, query, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch references: %w", err)
	}

	// Organize metrics and references by post ID
	metricsMap := make(map[int64]*PostPublicMetrics)
	for i := range metrics {
		metricsMap[metrics[i].PostID] = &metrics[i]
	}

	referencesMap := make(map[int64][]PostReference)
	for _, ref := range references {
		referencesMap[ref.SourcePostID] = append(referencesMap[ref.SourcePostID], ref)
	}

	// Combine everything
	result := make([]PostData, len(posts))
	for i, post := range posts {
		postData := PostData{
			Post:    post,
			Metrics: metricsMap[post.ID],
		}

		if len(referencesMap) > 0 {
			refs := referencesMap[post.ID]
			postData.References = &refs
		}

		result[i] = postData
	}

	return result, nextCursor, nil
}
