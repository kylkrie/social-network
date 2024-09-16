package postdb

import (
	"context"
	"fmt"

	"yabro.io/social-api/internal/util"
)

type UserPostInteraction struct {
	PostID       int64
	IsLiked      bool
	IsBookmarked bool
}

func (pdb *PostDB) GetUserPostInteractions(ctx context.Context, postIDs []int64, userID int64) (map[int64]UserPostInteraction, error) {
	query := `
        SELECT 
            p.id AS post_id,
            CASE WHEN l.user_id IS NOT NULL THEN true ELSE false END AS is_liked,

            CASE WHEN b.user_id IS NOT NULL THEN true ELSE false END AS is_bookmarked

        FROM unnest($1::bigint[]) p(id)
        LEFT JOIN post_likes l ON p.id = l.post_id AND l.user_id = $2
        LEFT JOIN post_bookmarks b ON p.id = b.post_id AND b.user_id = $2
        ORDER BY p.id
    `

	rows, err := pdb.db.QueryContext(ctx, query, postIDs, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user post interactions: %w", err)
	}

	defer rows.Close()

	var interactions []UserPostInteraction
	for rows.Next() {
		var interaction UserPostInteraction
		err := rows.Scan(&interaction.PostID, &interaction.IsLiked, &interaction.IsBookmarked)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user post interaction: %w", err)
		}
		interactions = append(interactions, interaction)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating user post interactions: %w", err)
	}

	return util.ArrToMap(interactions, func(i UserPostInteraction) int64 { return i.PostID }), nil
}
